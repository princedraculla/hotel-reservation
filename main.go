package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "server running properly")

	app := fiber.New()

	apiv1User := app.Group("api/v1/user")
	apiv1User.Get("", api.HandleGetUser)

	app.Listen(*listenAddr)
}
