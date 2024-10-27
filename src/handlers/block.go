package handlers

import (
	"go-eth/service"

	"github.com/gin-gonic/gin"
)

func GetLatestBlockHeight(c *gin.Context) {
	block, err := service.GetLatestBlockHeight()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"blockHeight": block})
}
