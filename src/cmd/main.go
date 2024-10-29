package main

import (
	"go-eth/api/route"
	"go-eth/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	db := app.Mongo.Database("go-eth")
	defer app.CloseMongoDBConnection()
	env := app.Env
	ethClient := app.EthClient

	gin := gin.Default()
	route.Setup(env, db, ethClient, gin)
	gin.Run(":8080")
}
