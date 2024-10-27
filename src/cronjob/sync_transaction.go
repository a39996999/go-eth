package cronjob

import (
	"context"
	"go-eth/consts"
	"go-eth/repositories"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron/v3"
)

func RunSyncTransaction() {
	client, err := ethclient.Dial(consts.CHAIN_RPC_URL)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
	}

	address := common.HexToAddress("0x2d259bfa2597Ac1218df4dF4603122d499530efE")
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Printf("Failed to get current block number: %v", err)
	}

	chainId := big.NewInt(consts.CHAIN_ID)

	c := cron.New()
	c.AddFunc("@every 30s", func() {
		startTime := time.Now()
		for blockNumber := 0; blockNumber <= int(currentBlock); blockNumber++ {
			block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
			if err != nil {
				log.Printf("Failed to get block: %v", err)
			}

			for _, tx := range block.Transactions() {
				msg, err := types.Sender(types.LatestSignerForChainID(chainId), tx)
				if err != nil {
					log.Printf("Failed to get transaction sender: %v", err)
					continue
				}
				from := msg.Hex()
				to := tx.To().Hex()
				if from == address.Hex() || to == address.Hex() {
					log.Printf("Transaction from: %v, to: %v, hash: %v", from, to, tx.Hash().Hex())
				}

				transaction := &repositories.Transaction{
					From:     from,
					To:       to,
					Hash:     tx.Hash().Hex(),
					Value:    tx.Value().String(),
					Gas:      tx.Gas(),
					GasPrice: tx.GasPrice().String(),
					Nonce:    tx.Nonce(),
					Data:     string(tx.Data()),
				}
				if _, err := transaction.CreateOne(); err != nil {
					log.Printf("Failed to create transaction: %v", err)
				}
			}
		}
		endTime := time.Now()
		log.Printf("Block number: %v", currentBlock)
		log.Printf("Time taken: %v", endTime.Sub(startTime))
	})
	c.Start()
}
