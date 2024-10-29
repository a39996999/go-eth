package route

import (
	"go-eth/api/controller"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func NewCurrencyRoute(ethClient *ethclient.Client, router *gin.RouterGroup) {
	currencyController := controller.CurrencyController{
		EthClient: ethClient,
	}
	router.POST("/currency/receive", currencyController.ReceiveNativeCoin)
}
