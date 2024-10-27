package cronjob

import (
	"go-eth/consts"
	"go-eth/repositories"
	"go-eth/service"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robfig/cron/v3"
)

func RunSyncTransaction() {
	c := cron.New()
	c.AddFunc("@every 15s", func() {
		users, err := (&repositories.User{}).GetAll()
		if err != nil {
			log.Printf("Failed to get users: %v", err)
			return
		}

		currentBlock, err := service.GetLatestBlockHeight()
		if err != nil {
			log.Printf("Failed to get current block number: %v", err)
			return
		}

		startTime := time.Now()
		user_map := make(map[string]*repositories.User)
		minBlockNumber := uint64(currentBlock)
		for _, user := range users {
			user_map[user.Address] = user
			if user.CurrentSyncBlock < minBlockNumber {
				minBlockNumber = user.CurrentSyncBlock
			}
		}

		blockBatchSize := 500
		for startBlock := minBlockNumber; startBlock <= uint64(currentBlock); startBlock += uint64(blockBatchSize) {
			endBlock := startBlock + uint64(blockBatchSize) - 1
			if endBlock > uint64(currentBlock) {
				endBlock = uint64(currentBlock)
			}

			go func(start, end uint64) {
				for blockNumber := start; blockNumber <= end; blockNumber++ {
					block, err := service.GetBlockByNumber(int64(blockNumber))
					if err != nil {
						log.Printf("Failed to get block: %v", err)
						continue
					}

					transactions := block.Transactions()
					for _, tx := range transactions {
						msg, err := types.Sender(types.LatestSignerForChainID(big.NewInt(consts.CHAIN_ID)), tx)
						if err != nil {
							log.Printf("Failed to get transaction sender: %v", err)
							continue
						}

						from := msg.Hex()
						to := tx.To().Hex()
						if user_map[from] != nil || user_map[to] != nil {
							log.Printf("Transaction from: %v, to: %v, hash: %v", from, to, tx.Hash().Hex())
						}

						transaction := &repositories.Transaction{
							From:      from,
							To:        to,
							Hash:      tx.Hash().Hex(),
							Value:     tx.Value().String(),
							Gas:       tx.Gas(),
							GasPrice:  tx.GasPrice().String(),
							Nonce:     tx.Nonce(),
							Data:      string(tx.Data()),
							Timestamp: int64(block.Time()),
						}
						if _, err := transaction.UpsertOne(); err != nil {
							log.Printf("Failed to upsert transaction: %v", err)
						}
					}
				}
			}(startBlock, endBlock)
		}

		for _, user := range users {
			user.CurrentSyncBlock = uint64(currentBlock)
			if _, err := user.UpdateBlockNumber(); err != nil {
				log.Printf("Failed to update user block number: %v", err)
			}
		}

		endTime := time.Now()
		log.Printf("Block number: %v", currentBlock)
		log.Printf("Time taken: %v", endTime.Sub(startTime))
	})
	c.Start()
}
