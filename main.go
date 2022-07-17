package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpparas/flight_log/database"
	"github.com/rpparas/flight_log/router"
)

func Setup() *fiber.App {
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)
	return app
}

func main() {
	// Start a new REST API app using fiber as framework
	app := Setup()
	app.Listen(":8000")
}
