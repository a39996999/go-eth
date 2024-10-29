package route

import (
	"go-eth/api/controller"
	"go-eth/bootstrap"
	"go-eth/domain"
	"go-eth/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRoute(env *bootstrap.Env, db *mongo.Database, router *gin.RouterGroup) {
	userRepository := repositories.NewUserRepository(db, domain.CollectionUser)
	userController := controller.UserController{
		UserRepository: userRepository,
		Env:            env,
	}
	router.POST("/user/create/:address", userController.CreateUser)
	router.GET("/user/balance/:address", userController.GetUserBalance)
}
