package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "server running properly")

	app := fiber.New()
	app.Listen(*listenAddr)
}
