package bootstrap

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env       *Env
	Mongo     *mongo.Client
	EthClient *ethclient.Client
}

func App() *Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	app.EthClient = NewEthClient(app.Env)
	return app
}

func (app *Application) CloseMongoDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
