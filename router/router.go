package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/rpparas/flight_log/docs"
	flightRoutes "github.com/rpparas/flight_log/internals/routes/flights"
	robotRoutes "github.com/rpparas/flight_log/internals/routes/robots"
)

func SetupRoutes(app *fiber.App) {
	// Group api calls with param '/api/v1'
	api := app.Group("/api/v1", logger.New())

	robotRoutes.SetupApiRoutes(api)
	flightRoutes.SetupApiRoutes(api)
}
