package domain

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CollectionUser = "users"
)

type User struct {
	Address          string `bson:"address"`
	CurrentSyncBlock int64  `bson:"current_sync_block"`
}

type UserRepository interface {
	UpsertOne(user *User) (*mongo.UpdateResult, error)
	GetUser(address string) (*User, error)
	UpdateBlockNumber(address string, blockNumber int64) (*mongo.UpdateResult, error)
	GetAll() ([]*User, error)
}
