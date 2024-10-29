package route

import (
	"go-eth/bootstrap"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, db *mongo.Database, ethClient *ethclient.Client, g *gin.Engine) {
	g.LoadHTMLGlob("./web/*")
	g.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	publicRouter := g.Group("")
	NewUserRoute(env, db, ethClient, publicRouter)
	NewBlockRoute(ethClient, publicRouter)
	NewTransactionRoute(env, db, publicRouter)
}
