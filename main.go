package main

import (
	"context"
	"flag"

	"github.com/gofiber/fiber/v2"
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
	listenAddr := flag.String("listenAddr", ":5000", "server running properly")

	//database connection
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	userHandler := api.NewUserHandler(db.NewMongoUserStorer(client))

	app := fiber.New(config)
	apiv1User := app.Group("api/v1/user")
	apiv1User.Get("/:id", userHandler.HandleGetUser)
	apiv1User.Get("/list", userHandler.HandleGetUsers)
	apiv1User.Post("/add", userHandler.HandlePostUser)

	app.Listen(*listenAddr)
}
