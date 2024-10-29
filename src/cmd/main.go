package main

import (
	"go-eth/api/route"
	"go-eth/bootstrap"
	"go-eth/cronjob"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	db := app.Mongo.Database("go-eth")
	defer app.CloseMongoDBConnection()
	defer app.CloseEthClient()
	env := app.Env
	ethClient := app.EthClient

	gin := gin.Default()
	route.Setup(env, db, ethClient, gin)
	cronjob.Setup(env, db, ethClient)
	gin.Run(":8080")
}
