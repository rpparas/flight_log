package flights

import (
	"github.com/gofiber/fiber/v2"
	flightsHandler "github.com/rpparas/flight_log/internals/handlers/flights"
)

func SetupApiRoutes(router fiber.Router) {
	// Endpoints for flights
	flights := router.Group("/flights")
	flights.Post("/", flightsHandler.CreateFlight)
	flights.Get("/", flightsHandler.GetFlights)
	flights.Get("/:flightId", flightsHandler.GetFlight)

}
