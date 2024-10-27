package cronjob

import (
	"context"
	"go-eth/consts"
	"go-eth/repositories"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron/v3"
)

func RunSyncTransaction() {
	c := cron.New()
	c.AddFunc("@every 15s", func() {
		client, err := ethclient.Dial(consts.CHAIN_RPC_URL)
		if err != nil {
			log.Printf("Failed to connect to the Ethereum client: %v", err)
		}

		users, err := (&repositories.User{}).GetAll()
		if err != nil {
			log.Printf("Failed to get users: %v", err)
		}

		currentBlock, err := client.BlockNumber(context.Background())
		if err != nil {
			log.Printf("Failed to get current block number: %v", err)
		}

		chainId := big.NewInt(consts.CHAIN_ID)
		startTime := time.Now()

		for _, user := range users {
			go func(user *repositories.User) {
				log.Printf("Syncing transactions for user: %v", user.Address)
				log.Printf("Current sync block: %v", user.CurrentSyncBlock)
				for blockNumber := user.CurrentSyncBlock; blockNumber <= currentBlock; blockNumber++ {
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
						if from == user.Address || to == user.Address {
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
						if _, err := transaction.UpsertOne(); err != nil {
							log.Printf("Failed to upsert transaction: %v", err)
						}
					}
				}
				user.CurrentSyncBlock = currentBlock
				if _, err := user.UpdateBlockNumber(); err != nil {
					log.Printf("Failed to update user block number: %v", err)
				}
			}(user)
		}

		endTime := time.Now()
		log.Printf("Block number: %v", currentBlock)
		log.Printf("Time taken: %v", endTime.Sub(startTime))
	})
	c.Start()
}
