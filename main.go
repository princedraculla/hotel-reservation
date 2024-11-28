package main

import (
	"context"
	"errors"
	"flag"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/api"
	"github.com/princedraculla/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri = "mongodb://localhost:27017"
var config = fiber.Config{

	ErrorHandler: func(ctx *fiber.Ctx, err error) error {

		code := fiber.StatusInternalServerError

		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}

		err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
		if err != nil {

			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return nil
	},
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

	app.Listen(*listenAddr)
}
