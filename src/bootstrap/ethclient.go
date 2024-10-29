package bootstrap

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthClient(env *Env) *ethclient.Client {
	ethClient, err := ethclient.Dial(env.RPCUrl)
	if err != nil {
		log.Fatalf("Error connecting to Ethereum client: %v", err)
	}
	return ethClient
}
