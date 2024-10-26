package router

import (
	"go-eth/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/block/latest", handlers.GetLatestBlockHeight)
	router.GET("/user/balance/:address", handlers.GetUserBalance)
	router.POST("/currency/send", handlers.SendNativeCoin)
	router.POST("/currency/receive", handlers.ReceiveNativeCoin)

	return router
}
