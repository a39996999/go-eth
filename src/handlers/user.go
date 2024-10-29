package handlers

import (
	"go-eth/service"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"go-eth/repositories"
)

func CreateUser(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		log.Println("Invalid address:", address)
		c.JSON(400, gin.H{"error": "Invalid address"})
		return
	}
	user := &repositories.User{Address: common.HexToAddress(address).Hex(), CurrentSyncBlock: 0}
	if _, err := user.UpsertOne(); err != nil {
		log.Println("Failed to upsert user:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User created"})
}

func GetUserBalance(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		log.Println("Invalid address:", address)
		c.JSON(400, gin.H{"error": "Invalid address"})
		return
	}

	balance, err := service.GetBalance(common.HexToAddress(address).Hex())
	if err != nil {
		log.Println("Failed to get balance:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ethBalance := decimal.NewFromBigInt(balance, 0)
	ethBalance = ethBalance.Div(decimal.NewFromInt(1e18)).Round(6)

	c.JSON(200, gin.H{"balance": ethBalance.String()})
}
