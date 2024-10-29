package main

import (
	"go-eth/api/route"
	"go-eth/bootstrap"
	"go-eth/service"

	"github.com/gin-gonic/gin"
)

func init() {
	service.InitEthClient()
}
func main() {
	app := bootstrap.App()

	db := app.Mongo.Database("go-eth")
	defer app.CloseMongoDBConnection()
	env := app.Env

	gin := gin.Default()
	route.Setup(env, db, gin)
	gin.Run(":8080")
}
