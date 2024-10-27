package service

import (
	"context"
	"go-eth/consts"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var EthClient *ethclient.Client

func InitEthClient() {
	var err error
	EthClient, err = ethclient.Dial(consts.CHAIN_RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
}

func GetLatestBlockHeight() (int64, error) {
	block, err := EthClient.BlockByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return block.Number().Int64(), nil
}

func GetPendingNonceAt(address string) (uint64, error) {
	nonce, err := EthClient.PendingNonceAt(context.Background(), common.HexToAddress(address))
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func SuggestGasPrice() (*big.Int, error) {
	gasPrice, err := EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}

func SendTransaction(tx *types.Transaction) error {
	err := EthClient.SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}
	return nil
}

func GetBalance(address string) (*big.Int, error) {
	balance, err := EthClient.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}
