package handlers

import (
	"go-eth/service"
	"log"

	"github.com/gin-gonic/gin"
)

func GetLatestBlockHeight(c *gin.Context) {
	block, err := service.GetLatestBlockHeight()
	if err != nil {
		log.Println("Failed to get latest block height:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"blockHeight": block})
}
