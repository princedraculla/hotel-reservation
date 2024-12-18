package db

import (
	"context"

	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(context.Context, bson.M, bson.M) error
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client:     client,
		collection: client.Database(DBNAME).Collection("hotels"),
	}
}

func (hotelStore *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	insertedHotel, err := hotelStore.collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = insertedHotel.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (hotelStore *MongoHotelStore) Update(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := hotelStore.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
