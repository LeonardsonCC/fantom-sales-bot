package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/machinebox/graphql"
)

const layoutISO = "02-01-2006"
const FANTOM_RPC_URL = "https://rpc.ftm.tools/"

type Transaction struct {
	BlockNumber     string `json:"blockNumber"`
	Hash            string `json:"hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	TimeStamp       string `json:"timeStamp"`
}

type CollectionData struct {
	Name    string `json:"name"`
	Twitter string `json:"twitter"`
}

type SaleHistoryItem struct {
	TxHash   string
	ActionId string
	Value    float64
	Time     int64
	Version  int32
	Token    string
}

type SalesHistory []SaleHistoryItem

func (a SalesHistory) Len() int {
	return len(a)
}
func (a SalesHistory) Swap(i, j int64) {
	a[i], a[j] = a[j], a[i]
}
func (a SalesHistory) Less(i, j int64) bool {
	return a[i].Time < a[j].Time
}

type SaleAction struct {
	Id        string
	Address   string
	NftId     string
	Seller    string
	MaxBidder string
}

type Sale struct {
	Id        string
	Addresses []string
	TokenIds  []string
	EndTime   int64
	Value     float64
	LastSales SalesHistory
}

type PriceIn struct {
	Usd float64 `json:"usd"`
	Brl float64 `json:"brl"`
}

type MarketData struct {
	CurrentPrice PriceIn `json:"current_price"`
}

type FantomPriceResponse struct {
	Symbol     string     `json:"symbol"`
	MarketData MarketData `json:"market_data"`
}

type TwitterUploadResponse struct {
	MediaId             int64  `json:"media_id"`
	MediaIdString       string `json:"media_id_string"`
	Size                int64  `json:"size"`
	ExpiresAfterSeconds int64  `json:"expires_after_secs"`
}

type NftApi struct {
	Image string `json:"image"`
}

type PaintSwapNft struct {
	Address string `json:"address"`
	TokenId uint   `json:"tokenId"`
	Uri     string `json:"uri"`
}

type PaintSwapNftResponse struct {
	Nft PaintSwapNft `json:"nft"`
}

var currentTime time.Time = time.Now()
var blocklist []string

func main() {
	fmt.Println("Starting...")

	blocklistLoaded, err := LoadBlocklist()
	if err != nil {
		fmt.Println("error loading blocklist")
		blocklist = []string{}
	}
	blocklist = blocklistLoaded

	for {
		twitterClient := getTwitterConfig()

		client := graphql.NewClient("https://api.thegraph.com/subgraphs/name/paint-swap-finance/nftsv2")

		recentSales := getRecentSales(client)
		for _, sale := range recentSales.([]interface{}) {
			address := sale.(map[string]interface{})["id"].(string)
			splittedAddres := strings.Split(address, "_")
			address = splittedAddres[0]

			tokenId := splittedAddres[1]

			go fetchSaleHistoryAndTweet(twitterClient, client, sale, tokenId, address)
		}
		fmt.Println("Waiting 5 minutes")
		time.Sleep(time.Minute * 5)
		currentTime = time.Now()
		fmt.Println("Less go")
	}
}

func getPrice(token string, dateString string) float64 {
	t, _ := time.Parse(layoutISO, dateString)
	validatedTime := t.Format(layoutISO)

	resp, err := http.Get(fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/history?date=%v", token, validatedTime))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response FantomPriceResponse
	json.Unmarshal(body, &response)

	if response.MarketData.CurrentPrice.Usd == 0 {
		fmt.Println("Coingecko returned 0 in USD price... Trying again")
		time.Sleep(time.Second * 10)
		return getPrice(token, dateString)
	}

	return response.MarketData.CurrentPrice.Usd
}

func fetchSaleHistoryAndTweet(twitterClient *twitter.Client, client *graphql.Client, sale interface{}, tokenId string, address string) {
	// Web3 Stuff
	conn, err := ethclient.Dial(FANTOM_RPC_URL)
	if err != nil {
		log.Fatal(err)
	}

	intTokenId, _ := strconv.ParseUint(tokenId, 10, 64)
	history := getSaleHistory(address, uint(intTokenId), client)
	if len(history.([]interface{})) > 0 {
		endTime, _ := strconv.ParseInt(sale.(map[string]interface{})["timestamp"].(string), 10, 64)
		price, _ := strconv.ParseFloat(sale.(map[string]interface{})["data"].(string), 64)
		var sale *Sale = &Sale{
			Id:        sale.(map[string]interface{})["id"].(string),
			Addresses: []string{address},
			TokenIds:  []string{tokenId},
			EndTime:   endTime,
			Value:     bigIntToLegibleNumber(price),
		}
		var salesHistory SalesHistory = make([]SaleHistoryItem, 0)

		for _, historyItem := range history.([]interface{}) {
			saleValue, _ := strconv.ParseFloat(historyItem.(map[string]interface{})["data"].(string), 32)
			saleTime, _ := strconv.ParseInt(historyItem.(map[string]interface{})["timestamp"].(string), 10, 64)
			version, _ := historyItem.(map[string]interface{})["version"].(float64)

			var oldSale *SaleHistoryItem = &SaleHistoryItem{
				TxHash:   historyItem.(map[string]interface{})["hash"].(string),
				ActionId: historyItem.(map[string]interface{})["actionId"].(string),
				Value:    bigIntToLegibleNumber(saleValue),
				Time:     saleTime,
				Version:  int32(version),
			}

			if oldSale.Version == 1 {
				oldSale.Token = "paint-swap"
			} else {
				oldSale.Token = "fantom"
			}

			salesHistory = append(salesHistory, *oldSale)
		}
		sale.LastSales = salesHistory

		contract, err := NewGenericContract(common.HexToAddress(address), conn)
		if err != nil {
			log.Fatal(err)
		}

		var boughtAction SaleHistoryItem
		var soldAction SaleHistoryItem
		var boughtOrMinted string
		if len(history.([]interface{})) > 1 {
			boughtOrMinted = "Bought"
			boughtAction = salesHistory[len(sale.LastSales)-2]
			soldAction = salesHistory[len(sale.LastSales)-1]
		} else {
			boughtOrMinted = "Minted"
			saleAction, err := GetSaleAction(client, address, tokenId)
			if err != nil {
				fmt.Println(err)
				return
			}

			oldOwner := saleAction.Seller
			transactions, _ := GetContractTxs(address)
			var mintTx Transaction
			var mintQty int64
			var foundIt bool = false
			for _, tx := range transactions {
				txValue, _ := strconv.ParseFloat(tx.Value, 10)
				if !foundIt && tx.From == oldOwner && tx.ContractAddress == "" && txValue > 0 {
					txWeb3, isPending, err := conn.TransactionByHash(context.Background(), common.HexToHash(tx.Hash))
					if err != nil {
						fmt.Println(err)
						return
					}

					if !isPending {
						receipt, err := conn.TransactionReceipt(context.Background(), txWeb3.Hash())
						if err != nil {
							fmt.Println(err)
							return
						}

						receiptJson, err := receipt.MarshalJSON()
						if err != nil {
							fmt.Println(err)
							break
						}

						var receiptMap struct {
							Logs []struct {
								Topics []string `json:"topics"`
							}
						}
						json.Unmarshal(receiptJson, &receiptMap)

						mintQty = 0
						for _, log := range receiptMap.Logs {
							if len(log.Topics) == 4 { // need 4 topics to make a transfer of ERC721
								mintQty = 1 + mintQty
								for _, topic := range log.Topics {
									num, _ := strconv.ParseInt(topic[2:], 16, 64)
									if strconv.Itoa(int(num)) == tokenId {
										mintTx = tx
										fmt.Println(tx)
										foundIt = true
										break
									}
								}
							}
						}
					}
				}
			}

			mintValue, _ := strconv.ParseFloat(mintTx.Value, 10)
			mintValue = mintValue / float64(mintQty)
			mintTimeStamp, _ := strconv.ParseInt(mintTx.TimeStamp, 10, 64)
			boughtAction = SaleHistoryItem{
				TxHash:   mintTx.Hash,
				ActionId: "0",
				Time:     mintTimeStamp,
				Value:    bigIntToLegibleNumber(mintValue),
				Token:    "fantom",
				Version:  2,
			}

			soldAction = salesHistory[len(sale.LastSales)-1]
			if math.IsNaN(boughtAction.Value) || boughtAction.Time == 0 {
				fmt.Printf("Contract skipped because some error in values: %+v\n", boughtAction)
				return
			}
		}

		if boughtAction.Token == "paint-swap" {
			brushPrice := getPrice(boughtAction.Token, time.Unix(boughtAction.Time, 0).Format(layoutISO))
			fantomPrice := getPrice("fantom", time.Unix(boughtAction.Time, 0).Format(layoutISO))

			totalBrushInUsd := boughtAction.Value * brushPrice
			totalFtmValue := totalBrushInUsd / fantomPrice

			boughtAction.Value = totalFtmValue
		}
		if soldAction.Token == "paint-swap" {
			brushPrice := getPrice(soldAction.Token, time.Unix(soldAction.Time, 0).Format(layoutISO))
			fantomPrice := getPrice("fantom", time.Unix(soldAction.Time, 0).Format(layoutISO))

			totalBrushInUsd := soldAction.Value * brushPrice
			totalFtmValue := totalBrushInUsd / fantomPrice

			soldAction.Value = totalFtmValue
		}

		fmt.Printf("%+v\n%+v\n", boughtAction, soldAction)

		var tweetMessage string = ""

		collectionData, err := getCollectionData(address)
		if err != nil {
			name, err := contract.Name(nil)
			if err != nil {
				tweetMessage += fmt.Sprintf("🧾 Collection: %v\n\n", strings.Join(sale.Addresses, ","))
			} else {
				tweetMessage += fmt.Sprintf("🧾 Collection: %v\n\n", name)
			}
		} else {
			tweetMessage += fmt.Sprintf("🧾 Collection: %s (@%s)\n\n", collectionData.Name, collectionData.Twitter)
		}

		boughtFantomPrice := getPrice("fantom", time.Unix(boughtAction.Time, 0).Format(layoutISO))
		soldFantomPrice := getPrice("fantom", time.Unix(soldAction.Time, 0).Format(layoutISO))

		tweetMessage += fmt.Sprintf("🛍 %s: %.2f FTM @ $%.2f\n", boughtOrMinted, boughtAction.Value, boughtFantomPrice)
		tweetMessage += fmt.Sprintf("💰 Sold: %.2f FTM @ $%.2f\n", soldAction.Value, soldFantomPrice)

		boughtAt := boughtAction.Value
		soldAt := soldAction.Value
		difference := soldAt - boughtAt
		differenceDollar := (soldAt * soldFantomPrice) - (boughtAt * boughtFantomPrice)

		dateDifference := time.Unix(soldAction.Time, 0).Sub(time.Unix(boughtAction.Time, 0))
		tweetMessage += fmt.Sprintf("🤝 HODL: %.0f days\n", dateDifference.Hours()/24)

		if difference > 0 {
			tweetMessage += fmt.Sprintf("📈 Gain: %.2f FTM\n", difference)
		} else {
			tweetMessage += fmt.Sprintf("📉 Loss: %.2f FTM\n", (difference * -1))
		}

		if differenceDollar > 0 {
			tweetMessage += fmt.Sprintf("💵 Profit: $%.2f (📈 %.2f%%)\n", differenceDollar, ((difference / boughtAt) * 100))
		} else {
			tweetMessage += fmt.Sprintf("💵 Loss: $%.2f (📉 %.2f%%)\n", (differenceDollar * -1), ((difference / boughtAt) * 100))
		}
		tweetMessage += fmt.Sprintf("https://paintswap.finance/marketplace/%v", soldAction.ActionId)

		fmt.Printf("====================\n%v\n====================\n", tweetMessage)

		var isBlocked bool = false
		for _, blocklistItem := range blocklist {
			if blocklistItem == address {
				isBlocked = true
			}
		}
		if isBlocked {
			fmt.Printf("Skipping %v\n", address)
		} else {
			Tweet(twitterClient, tweetMessage, address, tokenId)
		}
	}
}

func Tweet(twitterClient *twitter.Client, tweetMessage string, address string, tokenId string) {
	nftImage, err := getNftImage(address, tokenId)

	if err != nil {
		twitterClient.Statuses.Update(tweetMessage, &twitter.StatusUpdateParams{})
		fmt.Println("tweeting without image")
	} else {
		mediaId, err := uploadImageToTwitter(twitterClient, nftImage)
		if err != nil {
			twitterClient.Statuses.Update(tweetMessage, &twitter.StatusUpdateParams{})
			fmt.Println("tweeted without image")
		} else {
			twitterClient.Statuses.Update(tweetMessage, &twitter.StatusUpdateParams{
				MediaIds: []int64{mediaId},
			})
			fmt.Println("tweeted with image")
		}
	}
}

func GetContractTxs(address string) ([]Transaction, error) {
	apiKey := os.Getenv("FTMSCAN_API_KEY")
	resp, err := http.Get(fmt.Sprintf("https://api.ftmscan.com/api?module=account&action=txlist&address=%s&startblock=0&endblock=99999999&sort=asc&apikey=%s", address, apiKey))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type FtmScanResult struct {
		Result *[]Transaction `json:"result"`
	}
	var result FtmScanResult
	json.Unmarshal(body, &result)

	var transactions []Transaction = *result.Result
	return transactions, nil
}

func LoadBlocklist() ([]string, error) {
	data, err := os.ReadFile("./data/blocklist.json")
	if err != nil {
		return nil, errors.New("error reading blocks file")
	}

	var blocklist []string
	err = json.Unmarshal(data, &blocklist)
	if err != nil {
		return nil, errors.New("error parsing blocklist file json")
	}

	return blocklist, nil
}

func getCollectionData(address string) (*CollectionData, error) {
	data, err := os.ReadFile("./data/collections.json")
	if err != nil {
		return nil, errors.New("error reading collections' file")
	}

	var allCollectionData interface{}
	err = json.Unmarshal(data, &allCollectionData)
	if err != nil || allCollectionData.(map[string]interface{})[address] == nil {
		return nil, errors.New("error parsing collections file json")
	}

	var collectionData *CollectionData = &CollectionData{
		Name:    allCollectionData.(map[string]interface{})[address].(map[string]interface{})["name"].(string),
		Twitter: allCollectionData.(map[string]interface{})[address].(map[string]interface{})["twitter"].(string),
	}

	return collectionData, nil
}

func bigIntToLegibleNumber(bigInt float64) float64 {
	return math.Ceil(bigInt / math.Pow(10, 18))
}

func clearNftUrl(url string) string {
	return strings.Replace(url, "ipfs://", "https://cloudflare-ipfs.com/ipfs/", 1)
}

func getNftImage(contractAddress string, tokenId string) (string, error) {
	var nftUrl string = clearNftUrl(getNftImageUrl(contractAddress, tokenId))
	if !strings.HasPrefix(nftUrl, "https://") {
		return "", fmt.Errorf("invalid nft url: %v", nftUrl)
	}
	imageUrl := nftUrl

	resp, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("failed to get nft image url")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("failed to read nft image")
	}

	var response NftApi
	json.Unmarshal(body, &response)

	return response.Image, nil
}

func getNftImageUrl(contractAddress string, tokenId string) string {
	resp, err := http.Get(fmt.Sprintf("https://api.paintswap.finance/nft/%v/%v?allowNSFW=true&numToFetch=10", contractAddress, tokenId))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response PaintSwapNftResponse
	json.Unmarshal(body, &response)

	return response.Nft.Uri
}

func GetSaleAction(client *graphql.Client, address string, tokenId string) (*SaleAction, error) {
	query := fmt.Sprintf(`
		query {
			sales(
				first: 1
				where: {
					nftIds_contains: ["%s"]
				}
			) {
				id
				addresses
				nftIds
				seller
				maxBidder
			}
		}
	`, fmt.Sprintf("%s_%s", address, tokenId))

	req := graphql.NewRequest(query)
	ctx := context.Background()

	var response map[string]interface{}

	if err := client.Run(ctx, req, &response); err != nil {
		return nil, err
	}

	var sale *SaleAction = &SaleAction{
		Id:        response["sales"].([]interface{})[0].(map[string]interface{})["id"].(string),
		Address:   response["sales"].([]interface{})[0].(map[string]interface{})["addresses"].([]interface{})[0].(string),
		NftId:     response["sales"].([]interface{})[0].(map[string]interface{})["nftIds"].([]interface{})[0].(string),
		Seller:    response["sales"].([]interface{})[0].(map[string]interface{})["seller"].(string),
		MaxBidder: response["sales"].([]interface{})[0].(map[string]interface{})["maxBidder"].(string),
	}

	return sale, nil
}

func getSaleHistory(contractAddress string, tokenId uint, client *graphql.Client) interface{} {
	commonIndexes := make([]uint, 40)
	var idsToSearch []string
	for i := range commonIndexes {
		idsToSearch = append(idsToSearch, fmt.Sprintf("\"%s_%d_%v\"", contractAddress, tokenId, i))
	}

	query := fmt.Sprintf(`
		query {
			nfthistories(
        orderBy: timestamp
        orderDirection: asc
        where: {
					id_in: [%v]
					action: "Sold"
				}
      ) {
				id
				numericId
				actionId
				action
				data
				timestamp
				hash
				version
			}
		}
	`, strings.Join(idsToSearch, ","))

	req := graphql.NewRequest(query)
	ctx := context.Background()

	var response map[string]interface{}

	if err := client.Run(ctx, req, &response); err != nil {
		log.Fatal(err)
	}

	return response["nfthistories"]
}

func getRecentSales(client *graphql.Client) interface{} {
	currentTime = currentTime.Add(-(time.Minute * time.Duration(5)))

	query := fmt.Sprintf(`
		query {
			nfthistories(
			where: {
					timestamp_gte: %v,
					action: Sold
				}) {
					id
					action
					data
					hash
					actionId
					timestamp
					version
				}
			}
    `, currentTime.Unix())

	req := graphql.NewRequest(query)
	ctx := context.Background()

	var response map[string]interface{}

	if err := client.Run(ctx, req, &response); err != nil {
		log.Fatal(err)
	}

	return response["nfthistories"]
}

func uploadImageToTwitter(twitterClient *twitter.Client, nftUrl string) (int64, error) {
	respNft, errNft := http.Get(clearNftUrl(nftUrl))
	if errNft != nil {
		return 0, errors.New("failed to get nft image")
	}

	bodyNft, err := ioutil.ReadAll(respNft.Body)
	if err != nil {
		return 0, errors.New("failed to read nft image")
	}

	fmt.Println("Uploading image to twitter...")
	media, _, err := twitterClient.Media.Upload(bodyNft, "tweet_image")
	if err != nil {
		return 0, errors.New("failed to upload image to twitter")
	}

	return media.MediaID, nil
}

func getTwitterConfig() *twitter.Client {
	consumer_key, consumer_secret, access_token, access_secret := getTwitterCredentials()

	config := oauth1.NewConfig(consumer_key, consumer_secret)
	token := oauth1.NewToken(access_token, access_secret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	return client
}

func getTwitterCredentials() (string, string, string, string) {
	consumer_key := os.Getenv("CONSUMER_KEY")
	consumer_secret := os.Getenv("CONSUMER_SECRET")
	access_token := os.Getenv("ACCESS_TOKEN")
	access_secret := os.Getenv("ACCESS_SECRET")
	return consumer_key, consumer_secret, access_token, access_secret
}
