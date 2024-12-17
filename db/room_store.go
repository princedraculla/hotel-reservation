package db

import (
	"context"

	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoRoomStore(client *mongo.Client, dbname string) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		collection: client.Database(dbname).Collection("rooms"),
	}
}

func (roomStore *MongoRoomStore) InsertRoom(ctx context.Context, rooms *types.Room) (*types.Room, error) {
	insertedRoom, err := roomStore.collection.InsertOne(ctx, rooms)
	if err != nil {
		return nil, err
	}
	rooms.ID = insertedRoom.InsertedID.(primitive.ObjectID)
	return rooms, nil
}
