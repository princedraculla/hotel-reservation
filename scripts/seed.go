package main

import (
	"context"
	"fmt"
	"log"

	"github.com/princedraculla/hotel-reservation/db"
	"github.com/princedraculla/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(db.DBURI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Taj Mahal Hotel",
		Location: "Tehran,iran",
	}
	rooms := types.Room{
		Type:      types.DeluxeRoomType,
		BasePrice: 99.9,
	}

	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	rooms.HotelID = insertedHotel.ID

	insertedRoom, err := roomStore.InsertRoom(ctx, &rooms)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("seeding Database...!")
	fmt.Println(insertedHotel)
	fmt.Println(insertedRoom)

}
