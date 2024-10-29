package controller

import (
	"go-eth/bootstrap"
	"go-eth/domain"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionRepository domain.TransactionRepository
	Env                   *bootstrap.Env
}

func (t *TransactionController) GetAllTransactions(c *gin.Context) {
	walletAddress := c.Param("address")
	if !common.IsHexAddress(walletAddress) {
		log.Println("Invalid wallet address:", walletAddress)
		c.JSON(400, gin.H{"error": "Invalid wallet address"})
		return
	}

	transactions, err := t.TransactionRepository.GetUserTransactions(common.HexToAddress(walletAddress).Hex())
	if err != nil {
		log.Println("Failed to get user transactions:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(transactions) == 0 {
		log.Println("No transactions found for user:", walletAddress)
		c.JSON(200, []interface{}{})
		return
	}
	c.JSON(200, transactions)
}
