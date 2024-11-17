package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/princedraculla/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "server running properly")
	app := fiber.New()

	apiv1 := app.Group("api/v1")

	apiv1.Get("/amir", api.HandleGetUser)

	app.Listen(*listenAddr)
}
