package handlers

import "github.com/gin-gonic/gin"

func GetLatestBlockHeight(c *gin.Context) {
	// TODO: Implement the logic to get the latest block height
	c.JSON(200, gin.H{"blockHeight": "latest_block_height"})
}
