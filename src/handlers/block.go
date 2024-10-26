package handlers

import (
	"context"

	"go-eth/consts"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func GetLatestBlockHeight(c *gin.Context) {
	client, err := ethclient.Dial(consts.CHAIN_RPC_URL)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"blockHeight": block.Number()})
}
