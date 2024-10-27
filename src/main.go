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
	r.Run()
}
