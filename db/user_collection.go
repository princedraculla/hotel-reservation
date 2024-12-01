package db

import (
	"context"

	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBNAME string = "hotel-reservation"
const userColl string = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	UserList(context.Context) ([]*types.User, error)
	AddUser(context.Context, *types.User) (*types.User, error)
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

func (storer *MongoUserStore) AddUser(ctx context.Context, user *types.User) (*types.User, error) {
	insertedUser, err := storer.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = insertedUser.InsertedID.(primitive.ObjectID)
	return user, nil
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

func (storer *MongoUserStore) UserList(ctx context.Context) ([]*types.User, error) {
	findOptions := options.Find()
	cursor, err := storer.collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	var users []*types.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
