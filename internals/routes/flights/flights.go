package flights

import (
	"github.com/gofiber/fiber/v2"
	flightsHandler "github.com/rpparas/flight_log/internals/handlers/flights"
)

func SetupApiRoutes(router fiber.Router) {
	// TODO: use singular vs plural
	// for uploading flights in bulk using CSV
	flights := router.Group("/flights")
	flights.Post("/csv", flightsHandler.CreateFlights)
	flights.Get("/", flightsHandler.GetFlights)

	// for single flight
	flight := router.Group("/flight")
	flight.Post("/", flightsHandler.CreateFlight)
	flight.Get("/:flightId", flightsHandler.GetFlight)
}
