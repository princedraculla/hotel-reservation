package db

import (
	"context"

	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DBNAME string = "hotel-reservation"
const userColl string = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserStorer(client *mongo.Client) *MongoUserStore {

	return &MongoUserStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(userColl),
	}
}

func (storer *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if err := storer.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
