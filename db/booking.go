package db

import (
	"context"

	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BookingStore interface {
	BookingRoom(context.Context, *types.Booking) (*types.Booking, error)
}

type MongoBookingStore struct {
	client     *mongo.Client
	collection *mongo.Collection
	BookingStore
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client:     client,
		collection: client.Database(DBNAME).Collection("Booking"),
	}
}

func (b *MongoBookingStore) BookingRoom(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	result, err := b.collection.InsertOne(ctx, booking)
	if err != nil {
		return nil, err
	}
	booking.ID = result.InsertedID.(primitive.ObjectID)
	return booking, nil
}
