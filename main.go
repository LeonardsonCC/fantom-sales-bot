package main

import (
	"context"
	"encoding/json"
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

func main() {
	fmt.Println("Starting...")

	//twitterClient := getTwitterConfig()

	client := graphql.NewClient("https://api.thegraph.com/subgraphs/name/paint-swap-finance/nftsv2")

	recentSales := getRecentSales(client)
	for _, sale := range recentSales.([]interface{}) {
		address := sale.(map[string]interface{})["addresses"].([]interface{})[0].(string)
		tokenId := sale.(map[string]interface{})["tokenIds"].([]interface{})[0].(string)
		intTokenId, _ := strconv.ParseUint(tokenId, 10, 64)

		history := getSaleHistory(address, uint(intTokenId), client)
		if len(history.([]interface{})) > 1 {
			endTime, _ := strconv.ParseInt(sale.(map[string]interface{})["endTime"].(string), 10, 64)
			price, _ := strconv.ParseFloat(sale.(map[string]interface{})["price"].(string), 64)
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

				var oldSale *SaleHistoryItem = &SaleHistoryItem{
					TxHash:   historyItem.(map[string]interface{})["hash"].(string),
					ActionId: historyItem.(map[string]interface{})["actionId"].(string),
					Value:    bigIntToLegibleNumber(saleValue),
					Time:     saleTime,
				}
				salesHistory = append(salesHistory, *oldSale)
			}
			sale.LastSales = salesHistory

			var boughtAction SaleHistoryItem = salesHistory[len(sale.LastSales)-2]
			var soldAction SaleHistoryItem = salesHistory[len(sale.LastSales)-1]

			var tweetMessage string = ""

			boughtFantomPrice := getFantomPrice(time.Unix(boughtAction.Time, 0).Format(layoutISO))
			soldFantomPrice := getFantomPrice(time.Unix(soldAction.Time, 0).Format(layoutISO))

			tweetMessage += fmt.Sprintf("🛍 Bought: %v FTM @ $%.3f\n", boughtAction.Value, boughtFantomPrice)
			tweetMessage += fmt.Sprintf("💰 Sold: %v FTM @ $%.3f\n", soldAction.Value, soldFantomPrice)

			boughtAt := boughtAction.Value
			soldAt := soldAction.Value
			difference := soldAt - boughtAt
			differenceDollar := (soldAt * soldFantomPrice) - (boughtAt * boughtFantomPrice)

			if difference > 0 {
				tweetMessage += fmt.Sprintf("📈 Gain: %v FTM\n", difference)
			} else {
				tweetMessage += fmt.Sprintf("📉 Loss: %v FTM\n", (difference * -1))
			}

			if differenceDollar > 0 {
				tweetMessage += fmt.Sprintf("💵 Profit: $%.3f\n", differenceDollar)
			} else {
				tweetMessage += fmt.Sprintf("💵 Loss: $%.3f\n", (differenceDollar * -1))
			}

			tweetMessage += fmt.Sprintf("https://paintswap.finance/marketplace/%v", soldAction.ActionId)

			fmt.Println(tweetMessage + "\n")
		}
	}
}

func bigIntToLegibleNumber(bigInt float64) float64 {
	return math.Ceil(bigInt / math.Pow(10, 18))
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
	query := fmt.Sprintf(`
		query {
			sales(
				where: {
					isERC721s_contains:[true]
					endTime_gt: "%v"
					addresses: ["0xfd211f3b016a75bc8d73550ac5adc2f1cae780c0"]
          complete: true
          sold: true
				} 
				first: 100
				orderDirection: desc 
				orderBy: endTime
			) {
				id
				addresses
				tokenIds
				endTime
        price
			}
		}
	`, 1631365200)

	req := graphql.NewRequest(query)
	ctx := context.Background()

	var response map[string]interface{}

	if err := client.Run(ctx, req, &response); err != nil {
		log.Fatal(err)
	}

	return response["sales"]
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
