package cronjob

import (
	"go-eth/bootstrap"
	"go-eth/repositories"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, db *mongo.Database, ethClient *ethclient.Client) {
	userRepository := repositories.NewUserRepository(db)
	transactionRepository := repositories.NewTransactionRepository(db)
	syncTransaction := &SyncTransaction{
		EthClient:             ethClient,
		UserRepository:        userRepository,
		TransactionRepository: transactionRepository,
		Env:                   env,
	}
	syncTransaction.Run()
}
