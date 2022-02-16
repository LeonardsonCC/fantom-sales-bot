# Fantom NFT Sales Bot

This is the code of the [Twitter Bot](https://twitter.com/ftm_nftsalesbot) that posts about NFT sales
on [Fantom Network](https://fantom.foundation)!

## Why?

It started as a way for me to learn more about blockchain and NFT stuffs, but it turns out a big thing
and I don't have time to maintain the project, so I'm opening it.

## About the project

Ok, I know there's some really bad code here... But that's why I'm making this public!
The bot listens Fantom's JSON RPC to get events from [NFTKEY](https://nftkey.app/) and [PaintSwap](https://paintswap.finance/), then it is converted to a pattern inside application and used to post on Twitter and Discord.
If it don't have a previous sale, I'll look for the seller's transactions in the contract of the NFT,
then find the transaction that have a pattern of 3 parameters, like a transfer method from wallet 0x0 to the seller wallet and assume that is the mint transaction. Get the topics quantity of that transaction (usually is a topic to each token), and divide the quantity of topics by the value of the transaction.

## Marketplaces

- [NFTKEY](https://nftkey.app/)
- [PaintSwap](https://paintswap.finance/)

## Environment variables

| Variable         | Description                                    |
| ---------------- | ---------------------------------------------- |
| CONSUMER_KEY     | Twitter consumer key                           |
| CONSUMER_SECRET  | Twitter consumer secret                        |
| ACCESS_TOKEN     | Twitter access token                           |
| ACCESS_SECRET    | Twitter access secret                          |
| DISCORD_TOKEN    | Discord bot token                              |
| FTMSCAN_API_KEY  | API Key of [FTMScan](https://ftmscan.com/)     |
| MONGO_CONN_STING | MongoDB connection string (used in docker)     |
| COLLECTION       | (Deprecated) used to track only one collection |
| SHOW_GAINS       | (Deprecated) used to track only one collection |
| MINT_PRICE       | (Deprecated) used to track only one collection |

## How to run the project?

```sh
# clone the project
git clone git@github.com:LeonardsonCC/fantom-sales-bot.git

# add environment variables
cp .env.example .env
# and fill the variables

# run the project (using docker-compose)
docker-compose up --build -d
```
