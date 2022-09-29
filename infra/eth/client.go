package eth

import "github.com/ethereum/go-ethereum/ethclient"

type Client struct {
	c *ethclient.Client
}

func NewClient() (*Client, error) {
	client, err := ethclient.Dial("wss://wsapi.fantom.network/")
	if err != nil {
		return nil, err
	}

	return &Client{
		c: client,
	}, nil
}

func (c *Client) Client() *ethclient.Client {
	return c.c
}
