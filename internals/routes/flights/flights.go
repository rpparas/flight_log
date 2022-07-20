package flights

import (
	"github.com/gofiber/fiber/v2"
	flightsHandler "github.com/rpparas/flight_log/internals/handlers/flights"
)

func SetupApiRoutes(router fiber.Router) {
	flights := router.Group("/flights")
	// for uploading flights in bulk using CSV
	flights.Post("/csv", flightsHandler.CreateFlights)

	flights.Post("/", flightsHandler.CreateFlight)

	// for retrieving flights
	flights.Get("/", flightsHandler.GetFlights)
	flights.Get("/:flightId", flightsHandler.GetFlight)

	// TODO: provide endpoint to edit (patch) and delete existing flights

	// TODO: provide endpoint to create, read robots

}
