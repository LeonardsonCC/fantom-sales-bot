package main

import (
	"log"

	"github.com/LeonardsonCC/fantom-sales-bot/application/marketplaces"
	"github.com/LeonardsonCC/fantom-sales-bot/infra/eth"
)

func main() {
	log.Println("connecting to client...")
	c, err := eth.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("connecting to contracts...")
	nftkey := marketplaces.NewNftKey(c)
	paintswap := marketplaces.NewPaintSwap(c)

	watcher := marketplaces.NewWatcher(nftkey, paintswap)

	log.Println("subscribing to events...")
	sales := watcher.Subscribe()

	log.Println("waiting events...")
	for {
		select {
		case sale := <-sales:
			log.Printf("Sale: %+v\n", sale)
		}
	}
}
