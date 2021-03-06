package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rpparas/flight_log/database"
	"github.com/rpparas/flight_log/router"
)

func Setup() *fiber.App {
	app := fiber.New()

	// Setup the router
	router.SetupRoutes(app)
	return app
}

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	// Start a new REST API app using fiber as framework
	app := Setup()

	database.ConnectDB()

	app.Listen(":8000")
}
