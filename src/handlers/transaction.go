package handlers

import (
	"go-eth/repositories"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	walletAddress := c.Param("address")
	if !common.IsHexAddress(walletAddress) {
		log.Println("Invalid wallet address:", walletAddress)
		c.JSON(400, gin.H{"error": "Invalid wallet address"})
		return
	}
	transactions, err := (&repositories.Transaction{}).GetUserTransactions(common.HexToAddress(walletAddress).Hex())
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
