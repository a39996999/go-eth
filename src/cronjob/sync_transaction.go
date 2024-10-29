package cronjob

import (
	"context"
	"go-eth/bootstrap"
	"go-eth/domain"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron/v3"
)

type SyncTransaction struct {
	EthClient             *ethclient.Client
	UserRepository        domain.UserRepository
	TransactionRepository domain.TransactionRepository
	Env                   *bootstrap.Env
}

func (s *SyncTransaction) Run() {
	c := cron.New()
	c.AddFunc("@every 15s", func() {
		users, err := s.UserRepository.GetAll()
		if err != nil {
			log.Printf("Failed to get users: %v", err)
			return
		}

		currentBlock, err := s.EthClient.BlockNumber(context.Background())
		if err != nil {
			log.Printf("Failed to get current block number: %v", err)
			return
		}

		startTime := time.Now()
		user_map := make(map[string]*domain.User)
		minBlockNumber := currentBlock
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
					block, err := s.EthClient.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
					if err != nil {
						log.Printf("Failed to get block: %v", err)
						continue
					}

					transactions := block.Transactions()
					for _, tx := range transactions {
						msg, err := types.Sender(types.LatestSignerForChainID(big.NewInt(int64(s.Env.ChainID))), tx)
						if err != nil {
							log.Printf("Failed to get transaction sender: %v", err)
							continue
						}

						from := msg.Hex()
						to := tx.To().Hex()
						if user_map[from] != nil || user_map[to] != nil {
							log.Printf("Transaction from: %v, to: %v, hash: %v", from, to, tx.Hash().Hex())
						}

						transaction := &domain.Transaction{
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
						if _, err := s.TransactionRepository.UpsertTransaction(transaction); err != nil {
							log.Printf("Failed to upsert transaction: %v", err)
						}
					}
				}
			}(startBlock, endBlock)
		}

		for _, user := range users {
			if _, err := s.UserRepository.UpdateBlockNumber(user.Address, currentBlock); err != nil {
				log.Printf("Failed to update user block number: %v", err)
			}
		}

		endTime := time.Now()
		log.Printf("Block number: %v", currentBlock)
		log.Printf("Time taken: %v", endTime.Sub(startTime))
	})
	c.Start()
}
