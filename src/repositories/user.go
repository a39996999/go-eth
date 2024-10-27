package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Address          string `bson:"address"`
	CurrentSyncBlock uint64 `bson:"current_sync_block,omitempty"`
}

func (user *User) UpsertOne() (*mongo.UpdateResult, error) {
	return client.Database("go-eth").
		Collection("users").
		UpdateOne(context.TODO(), bson.M{"address": user.Address},
			bson.M{"$set": bson.M{"address": user.Address}},
			options.Update().SetUpsert(true),
		)
}

func (user *User) GetOne() (*User, error) {
	err := client.Database("go-eth").
		Collection("users").
		FindOne(context.TODO(), bson.M{"address": user.Address}).Decode(user)
	return user, err
}

func (user *User) UpdateBlockNumber() (*mongo.UpdateResult, error) {
	return client.Database("go-eth").
		Collection("users").
		UpdateOne(context.TODO(), bson.M{"address": user.Address},
			bson.M{"$set": bson.M{"current_sync_block": user.CurrentSyncBlock}})
}

func (user *User) GetAll() ([]*User, error) {
	cursor, err := client.Database("go-eth").
		Collection("users").
		Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var users []*User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
