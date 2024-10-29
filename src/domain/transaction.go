package domain

import "go.mongodb.org/mongo-driver/mongo"

const (
	CollectionTransaction = "transactions"
)

type Transaction struct {
	From      string `bson:"from"`
	To        string `bson:"to"`
	Hash      string `bson:"hash"`
	Value     string `bson:"value"`
	Gas       uint64 `bson:"gas"`
	GasPrice  string `bson:"gas_price"`
	Nonce     uint64 `bson:"nonce"`
	Data      string `bson:"data"`
	Timestamp int64  `bson:"timestamp"`
}

type TransactionRepository interface {
	UpsertTransaction(transaction *Transaction) (*mongo.UpdateResult, error)
	GetUserTransactions(address string) ([]*Transaction, error)
}
