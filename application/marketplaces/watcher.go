package marketplaces

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Watcher struct {
	marketplaces []IMarketplace
}

type Sale struct {
	Address common.Address
	TokenId *big.Int
	Seller  common.Address
	Buyer   common.Address
	Value   *big.Int
}

func NewWatcher(marketplaces ...IMarketplace) *Watcher {
	return &Watcher{
		marketplaces: marketplaces,
	}
}

func (w *Watcher) Subscribe() chan *Sale {
	sales := make(chan *Sale)

	for _, mp := range w.marketplaces {
		err := mp.Subscribe(sales)
		if err != nil {
			log.Println("ERROR", err)
		}
	}

	return sales
}
