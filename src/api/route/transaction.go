package route

import (
	"go-eth/api/controller"
	"go-eth/bootstrap"
	"go-eth/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTransactionRoute(env *bootstrap.Env, db *mongo.Database, router *gin.RouterGroup) {
	transactionRepository := repositories.NewTransactionRepository(db)
	transactionController := controller.TransactionController{
		TransactionRepository: transactionRepository,
		Env:                   env,
	}
	router.GET("/transactions/:address", transactionController.GetAllTransactions)
}
