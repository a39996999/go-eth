package route

import (
	"go-eth/bootstrap"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, db *mongo.Database, g *gin.Engine) {
	g.LoadHTMLGlob("./web/*")
	g.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	publicRouter := g.Group("")
	NewUserRoute(env, db, publicRouter)
}
