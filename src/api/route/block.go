package route

import (
	"go-eth/api/controller"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func NewBlockRoute(ethClient *ethclient.Client, router *gin.RouterGroup) {
	blockController := controller.BlockController{
		EthClient: ethClient,
	}
	router.GET("/block/latest", blockController.GetLatestBlockHeight)
}
