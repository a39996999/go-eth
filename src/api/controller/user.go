package controller

import (
	"context"
	"go-eth/domain"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"go-eth/bootstrap"
)

type UserController struct {
	UserRepository domain.UserRepository
	EthClient      *ethclient.Client
	Env            *bootstrap.Env
}

func (uc *UserController) CreateUser(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		log.Println("Invalid address:", address)
		c.JSON(400, gin.H{"error": "Invalid address"})
		return
	}
	user := &domain.User{Address: common.HexToAddress(address).Hex(), CurrentSyncBlock: 0}
	if _, err := uc.UserRepository.UpsertOne(user); err != nil {
		log.Println("Failed to upsert user:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User created"})
}

func (uc *UserController) GetUserBalance(c *gin.Context) {
	address := c.Param("address")
	if !common.IsHexAddress(address) {
		log.Println("Invalid address:", address)
		c.JSON(400, gin.H{"error": "Invalid address"})
		return
	}

	balance, err := uc.EthClient.BalanceAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		log.Println("Failed to get balance:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ethBalance := decimal.NewFromBigInt(balance, 0)
	ethBalance = ethBalance.Div(decimal.NewFromInt(1e18)).Round(6)

	c.JSON(200, gin.H{"balance": ethBalance.String()})
}
