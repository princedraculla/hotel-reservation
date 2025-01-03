package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/princedraculla/hotel-reservation/api"
	"github.com/princedraculla/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb://localhost:27017"
var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("error for loading .env file %s", err)
	}
	listenAddr := flag.String("listenAddr", ":5000", "server running properly")

	//database connection
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	//room && hotel store init
	var (
		hotelStore = db.NewMongoHotelStore(client)
		roomStore  = db.NewMongoRoomStore(client, hotelStore)
		userStore  = db.NewMongoUserStorer(client, db.DBNAME)
		store      = &db.Store{
			User:  userStore,
			Hotel: hotelStore,
			Rooms: roomStore,
		}
	)
	// handlers
	var (
		app          = fiber.New(config)
		hotelHandler = api.NewHotelHandler(store)
		userHandler  = api.NewUserHandler(userStore)
		authHandler  = api.NewAuthHandler(userStore)
		apiv1User    = app.Group("api/v1/user", api.JWTAuthenticate(userStore))
	)

	// register and login
	app.Post("/api/register", authHandler.HandleAutheticate)

	// user APIs
	apiv1User.Put("/:id", userHandler.HandleUserUpdate)
	apiv1User.Get("/:id", userHandler.HandleGetUser)
	apiv1User.Post("/add", userHandler.HandlePostUser)
	apiv1User.Delete("/:id", userHandler.HandleDeleteUser)
	app.Get("/list", userHandler.HandleGetUsers)

	// hotel APIs
	app.Get("/hotels", hotelHandler.GetHotels)
	app.Get("/hotels/:id/rooms", hotelHandler.GetRooms)

	//server start
	app.Listen(*listenAddr)
}
