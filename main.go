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
			fmt.Printf("Valor: %.2f\n", (math.Ceil(saleValue / math.Pow(10, 18))))
		}
	}
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
	// fmt.Println(query)

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
					addresses: ["0xd5eb80f437c318b3bf8b3af985224966a3054f76"]
				} 
				first: 10
				orderDirection: desc 
				orderBy: endTime
			) {
				id
				addresses
				tokenIds
				endTime
			}
		}
	`, 1637623400)
	fmt.Println(query)

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
