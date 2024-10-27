package main

import (
	"go-eth/cronjob"
	"go-eth/repositories"
	"go-eth/router"
	"go-eth/service"
)

func init() {
	repositories.InitConnection()
	service.InitEthClient()
}

func main() {
	cronjob.RunSyncTransaction()
	r := router.InitRouter()
	r.Run(":8080")
}
