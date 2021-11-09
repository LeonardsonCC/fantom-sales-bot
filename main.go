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
	"github.com/machinebox/graphql"
)

const layoutISO = "02-01-2006"

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

func main() {
	fmt.Println("Starting...")

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
	intTokenId, _ := strconv.ParseUint(tokenId, 10, 64)
	history := getSaleHistory(address, uint(intTokenId), client)
	if len(history.([]interface{})) > 1 {
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
				brushPrice := getPrice(oldSale.Token, time.Unix(oldSale.Time, 0).Format(layoutISO))
				fantomPrice := getPrice("fantom", time.Unix(oldSale.Time, 0).Format(layoutISO))

				totalBrushInUsd := oldSale.Value * brushPrice
				totalFtmValue := totalBrushInUsd / fantomPrice

				oldSale.Value = totalFtmValue
			} else {
				oldSale.Token = "fantom"
			}

			salesHistory = append(salesHistory, *oldSale)
		}
		sale.LastSales = salesHistory

		var boughtAction SaleHistoryItem = salesHistory[len(sale.LastSales)-2]
		var soldAction SaleHistoryItem = salesHistory[len(sale.LastSales)-1]

		var tweetMessage string = ""

		tweetMessage += fmt.Sprintf("🧾 Contract: %v\n\n", strings.Join(sale.Addresses, ","))

		boughtFantomPrice := getPrice("fantom", time.Unix(boughtAction.Time, 0).Format(layoutISO))
		soldFantomPrice := getPrice("fantom", time.Unix(soldAction.Time, 0).Format(layoutISO))

		tweetMessage += fmt.Sprintf("🛍 Bought: %.2f FTM @ $%.2f\n", boughtAction.Value, boughtFantomPrice)
		tweetMessage += fmt.Sprintf("💰 Sold: %.2f FTM @ $%.2f\n", soldAction.Value, soldFantomPrice)

		boughtAt := boughtAction.Value
		soldAt := soldAction.Value
		difference := soldAt - boughtAt
		differenceDollar := (soldAt * soldFantomPrice) - (boughtAt * boughtFantomPrice)

		dateDifference := time.Unix(soldAction.Time, 0).Sub(time.Unix(boughtAction.Time, 0))
		tweetMessage += fmt.Sprintf("🤝 HODL duration: %.0f days\n", dateDifference.Hours()/24)

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

		fmt.Println(tweetMessage + "\n")
		nftImage, err := getNftImage(address, tokenId)
		if err != nil {
			fmt.Println("error getting nft image")
			twitterClient.Statuses.Update(tweetMessage, &twitter.StatusUpdateParams{})
		} else {
			mediaId := uploadImageToTwitter(twitterClient, nftImage)
			twitterClient.Statuses.Update(tweetMessage, &twitter.StatusUpdateParams{
				MediaIds: []int64{mediaId},
			})
		}
	}
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
		return "", errors.New(fmt.Sprintf("invalid nft url: %v", nftUrl))
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

func getFantomPrice(dateString string) float64 {
	t, _ := time.Parse(layoutISO, dateString)
	validatedTime := t.Format(layoutISO)

	resp, err := http.Get(fmt.Sprintf("https://api.coingecko.com/api/v3/coins/fantom/history?date=%v", validatedTime))
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
		return getFantomPrice(dateString)
	}

	return response.MarketData.CurrentPrice.Usd
}

func getSaleHistory(contractAddress string, tokenId uint, client *graphql.Client) interface{} {
	commonIndexes := make([]uint, 10)
	var idsToSearch []string
	for i := range commonIndexes {
		idsToSearch = append(idsToSearch, fmt.Sprintf("\"%s_%d_%v\"", contractAddress, tokenId, i))
	}

	query := fmt.Sprintf(`
		query {
			nfthistories(where: {
				id_in: [%v]
				action: "Sold"
			}) {
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

func uploadImageToTwitter(twitterClient *twitter.Client, nftUrl string) int64 {
	respNft, errNft := http.Get(clearNftUrl(nftUrl))
	if errNft != nil {
		fmt.Println("error upload_image")
		time.Sleep(time.Second * 10)
		uploadImageToTwitter(twitterClient, nftUrl)
	}

	bodyNft, err := ioutil.ReadAll(respNft.Body)
	if err != nil {
		fmt.Println("error upload_image reading body")
		time.Sleep(time.Second * 10)
		uploadImageToTwitter(twitterClient, nftUrl)
	}

	fmt.Println("Uploading image to twitter...")
	media, _, err := twitterClient.Media.Upload(bodyNft, "tweet_image")
	if err != nil {
		fmt.Println("error upload_image request")
		time.Sleep(time.Second * 10)
		uploadImageToTwitter(twitterClient, nftUrl)
	}

	return media.MediaID
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
