package router

import (
	"go-eth/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("./web/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})
	router.GET("/block/latest", handlers.GetLatestBlockHeight)
	router.GET("/user/balance/:address", handlers.GetUserBalance)
	router.POST("/user/create/:address", handlers.CreateUser)
	router.POST("/currency/send", handlers.SendNativeCoin)
	router.POST("/currency/receive", handlers.ReceiveNativeCoin)

	return router
}
