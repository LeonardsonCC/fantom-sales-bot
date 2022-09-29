package marketplaces

import (
	"log"

	"github.com/LeonardsonCC/fantom-sales-bot/contracts"
	"github.com/LeonardsonCC/fantom-sales-bot/infra/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type NftKey struct {
	Instance *contracts.Nftkey
}

func NewNftKey(client *eth.Client) *NftKey {
	var address = common.HexToAddress("0x1a7d6ed890b6c284271ad27e7abe8fb5211d0739")

	i, err := contracts.NewNftkey(address, client.Client())
	if err != nil {
		log.Fatal(err)
	}

	return &NftKey{
		Instance: i,
	}
}

func (n *NftKey) Subscribe(sales chan *Sale) error {
	tokenBoughts := make(chan *contracts.NftkeyTokenBought)

	_, err := n.Instance.WatchTokenBought(&bind.WatchOpts{}, tokenBoughts, nil, nil, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case e := <-tokenBoughts:
				s := &Sale{
					Address: e.Erc721Address,
					TokenId: e.TokenId,
					Seller:  e.Listing.Seller,
					Buyer:   e.Buyer,
					Value:   e.Listing.Value,
				}
				sales <- s
			}
		}
	}()

	return nil
}
