package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
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

func (t *Transaction) UpsertOne() (*mongo.UpdateResult, error) {
	result, err := client.Database("go-eth").
		Collection("transactions").
		UpdateOne(context.TODO(),
			bson.M{"hash": t.Hash},
			bson.M{"$set": t},
			options.Update().SetUpsert(true),
		)
	return result, err
}

func (t *Transaction) GetUserTransactions(address string) ([]*Transaction, error) {
	cursor, err := client.Database("go-eth").
		Collection("transactions").
		Find(context.TODO(), bson.M{"$or": []bson.M{
			{"from": address},
			{"to": address},
		}}, options.Find().SetSort(bson.M{"timestamp": -1}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var transactions []*Transaction
	if err := cursor.All(context.TODO(), &transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}
