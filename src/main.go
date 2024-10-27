package main

import (
	"go-eth/cronjob"
	"go-eth/repositories"
	"go-eth/router"
)

func init() {
	repositories.InitConnection()
}

func main() {
	cronjob.RunSyncTransaction()
	r := router.InitRouter()
	r.Run("127.0.0.1:8080")
}
