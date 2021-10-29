package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/machinebox/graphql"
)

type Sales []struct {
	Id        string
	Addresses []string
	TokenIds  []string
}

func main() {
	fmt.Println("Starting...")

	//twitterClient := getTwitterConfig()

	client := graphql.NewClient("https://api.thegraph.com/subgraphs/name/paint-swap-finance/nftsv2")

	recentSales := getRecentSales(client)
	for _, sale := range recentSales.([]interface{}) {
		fmt.Println(fmt.Sprintf("%v", sale))

		address := sale.(map[string]interface{})["addresses"].([]interface{})[0].(string)
		tokenId := sale.(map[string]interface{})["tokenIds"].([]interface{})[0].(string)
		intTokenId, _ := strconv.ParseUint(tokenId, 10, 64)

		history := getSaleHistory(address, uint(intTokenId), client)
		for _, historyItem := range history.([]interface{}) {
			fmt.Println(fmt.Sprintf("%v", historyItem))
			saleValue, _ := strconv.ParseFloat(historyItem.(map[string]interface{})["data"].(string), 32)
			fmt.Printf("Valor: %.2f\n", bigIntToLegibleNumber(saleValue))
		}
	}
}

func bigIntToLegibleNumber(bigInt float64) float64 {
	return math.Ceil(bigInt / math.Pow(10, 18))
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
				} 
				first: 100
				orderDirection: desc 
				orderBy: endTime
			) {
				id
				addresses
				tokenIds
				endTime
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
