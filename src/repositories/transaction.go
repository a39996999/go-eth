package repositories

import (
	"context"
	"go-eth/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type transactionRepository struct {
	client     *mongo.Database
	collection string
}

func NewTransactionRepository(client *mongo.Database) domain.TransactionRepository {
	return &transactionRepository{
		client:     client,
		collection: domain.CollectionTransaction,
	}
}

func (t *transactionRepository) UpsertTransaction(transaction *domain.Transaction) (*mongo.UpdateResult, error) {
	collection := t.client.Collection(t.collection)
	result, err := collection.
		UpdateOne(context.TODO(),
			bson.M{"hash": transaction.Hash},
			bson.M{"$set": transaction},
			options.Update().SetUpsert(true),
		)
	return result, err
}

func (t *transactionRepository) GetUserTransactions(address string) ([]*domain.Transaction, error) {
	collection := t.client.Collection(t.collection)
	cursor, err := collection.
		Find(context.TODO(), bson.M{"$or": []bson.M{
			{"from": address},
			{"to": address},
		}}, options.Find().SetSort(bson.M{"timestamp": -1}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var transactions []*domain.Transaction
	if err := cursor.All(context.TODO(), &transactions); err != nil {
		return nil, err
	}
	return transactions, nil
}
