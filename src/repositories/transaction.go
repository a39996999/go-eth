package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	From     string `bson:"from"`
	To       string `bson:"to"`
	Hash     string `bson:"hash"`
	Value    string `bson:"value"`
	Gas      uint64 `bson:"gas"`
	GasPrice string `bson:"gas_price"`
	Nonce    uint64 `bson:"nonce"`
	Data     string `bson:"data"`
}

func (t *Transaction) CreateOne() (*mongo.UpdateResult, error) {
	result, err := client.Database("go-eth").
		Collection("transactions").
		UpdateOne(context.TODO(),
			bson.M{"hash": t.Hash},
			bson.M{"$set": t},
			options.Update().SetUpsert(true),
		)
	return result, err
}
