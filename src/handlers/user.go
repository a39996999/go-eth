package handlers

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"go-eth/consts"
)

func GetUserBalance(c *gin.Context) {
	address := common.HexToAddress(c.Param("address"))

	client, err := ethclient.Dial(consts.CHAIN_RPC_URL)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	balance, err := client.BalanceAt(context.Background(), address, nil)
	fmt.Println(balance.String())
	var ethBalance big.Float
	ethBalance.SetString(balance.String())
	ethBalance.Quo(&ethBalance, big.NewFloat(1e18)).Text('g', 6)
	fmt.Println(ethBalance.String())

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"balance": ethBalance.String()})
}
