package flights

import (
	"github.com/gofiber/fiber/v2"
	flightsHandler "github.com/rpparas/flight_log/internals/handlers/flights"
)

func SetupApiRoutes(router fiber.Router) {
	flights := router.Group("/flights")

	// Create a single Flight
	flights.Post("/", flightsHandler.CreateFlight)

	// Read all Flights
	flights.Get("/", flightsHandler.GetFlights)

	// Read single Flight
	flights.Get("/:flightId", flightsHandler.GetFlight)
}
