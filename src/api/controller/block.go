package controller

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type BlockController struct {
	EthClient *ethclient.Client
}

func (bc *BlockController) GetLatestBlockHeight(c *gin.Context) {
	block, err := bc.EthClient.BlockByNumber(context.Background(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"blockHeight": block.Number().Int64()})
}
