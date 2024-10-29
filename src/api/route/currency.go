package route

import (
	"go-eth/api/controller"
	"go-eth/bootstrap"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func NewCurrencyRoute(env *bootstrap.Env, ethClient *ethclient.Client, router *gin.RouterGroup) {
	currencyController := controller.CurrencyController{
		EthClient: ethClient,
		Env:       env,
	}
	router.POST("/currency/receive", currencyController.ReceiveNativeCoin)
}
