package db

import (
	"context"
	"fmt"

	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoomStore interface {
	InsertRoom(context.Context, *types.Room) (*types.Room, error)
	GetRooms(context.Context, bson.M) ([]*types.Room, error)
}

type MongoRoomStore struct {
	client     *mongo.Client
	collection *mongo.Collection
	HotelStore
}

func NewMongoRoomStore(client *mongo.Client, hotelStore HotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:     client,
		collection: client.Database(DBNAME).Collection("rooms"),
		HotelStore: hotelStore,
	}
}

func (roomStore *MongoRoomStore) InsertRoom(ctx context.Context, rooms *types.Room) (*types.Room, error) {
	insertedRoom, err := roomStore.collection.InsertOne(ctx, rooms)
	if err != nil {
		return nil, err
	}
	rooms.ID = insertedRoom.InsertedID.(primitive.ObjectID)
	filter := bson.M{"_id": rooms.HotelID}
	update := bson.M{"$push": bson.M{"rooms": rooms.ID}}
	if err = roomStore.HotelStore.Update(ctx, filter, update); err != nil {
		return nil, err
	}
	fmt.Println(rooms)
	return rooms, nil
}

func (roomStore *MongoRoomStore) GetRooms(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	result, err := roomStore.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var rooms []*types.Room
	if err := result.All(ctx, &rooms); err != nil {
		return nil, err
	}
	return rooms, nil
}
