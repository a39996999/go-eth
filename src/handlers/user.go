package handlers

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"go-eth/consts"
	"go-eth/repositories"
)

func CreateUser(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		c.JSON(400, gin.H{"error": "Invalid address"})
		return
	}
	user := &repositories.User{Address: address, CurrentSyncBlock: 0}
	if _, err := user.UpsertOne(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User created"})
}

func GetUserBalance(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		c.JSON(400, gin.H{"error": "Invalid address"})
		return
	}

	client, err := ethclient.Dial(consts.CHAIN_RPC_URL)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	ethBalance := decimal.NewFromBigInt(balance, 0)
	ethBalance = ethBalance.Div(decimal.NewFromInt(1e18)).Round(6)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"balance": ethBalance.String()})
}
