package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rpparas/flight_log/database"
	"github.com/rpparas/flight_log/router"
)

func main() {
	// Start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the router
	router.SetupRoutes(app)

	// Listen on PORT 3000
	app.Listen(":8000")
}
