package repositories

import (
	"context"
	"go-eth/domain"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type userRepository struct {
	db         *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &userRepository{db, collection}
}

func (ur *userRepository) UpsertOne(user *domain.User) (*mongo.UpdateResult, error) {
	collection := ur.db.Collection(ur.collection)

	result, err := collection.UpdateOne(context.TODO(), bson.M{"address": user.Address},
		bson.M{"$set": bson.M{"current_sync_block": user.CurrentSyncBlock}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ur *userRepository) GetUser(address string) (*domain.User, error) {
	collection := ur.db.Collection(ur.collection)
	result := collection.FindOne(context.TODO(), bson.M{"address": address})
	if result.Err() != nil {
		return nil, result.Err()
	}
	var user domain.User
	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) UpdateBlockNumber(address string, blockNumber int64) (*mongo.UpdateResult, error) {
	collection := ur.db.Collection(ur.collection)

	result, err := collection.UpdateOne(context.TODO(), bson.M{"address": address},
		bson.M{"$set": bson.M{"current_sync_block": blockNumber}})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (ur *userRepository) GetAll() ([]*domain.User, error) {
	cursor, err := ur.db.Collection(ur.collection).Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []*domain.User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
