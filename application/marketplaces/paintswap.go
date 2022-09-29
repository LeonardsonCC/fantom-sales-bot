package marketplaces

import (
	"log"

	"github.com/LeonardsonCC/fantom-sales-bot/contracts"
	"github.com/LeonardsonCC/fantom-sales-bot/infra/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type PaintSwap struct {
	Instance *contracts.Paintswap
}

func NewPaintSwap(client *eth.Client) *PaintSwap {
	var address = common.HexToAddress("0xf3dF7b6DCcC267393784a3876d0CbCBDC73147d4")

	i, err := contracts.NewPaintswap(address, client.Client())
	if err != nil {
		log.Fatal(err)
	}

	return &PaintSwap{
		Instance: i,
	}
}

func (n *PaintSwap) Subscribe(sales chan *Sale) error {
	tokenSold := make(chan *contracts.PaintswapSold)

	_, err := n.Instance.WatchSold(&bind.WatchOpts{}, tokenSold)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case e := <-tokenSold:
				// for each token id in event
				for _, token := range e.TokenIds {
					s := &Sale{
						Address: e.Nfts[0],
						TokenId: token,
						Seller:  e.Seller,
						Buyer:   e.Buyer,
						Value:   e.Price,
					}
					sales <- s
				}
			}
		}
	}()

	return nil
}
