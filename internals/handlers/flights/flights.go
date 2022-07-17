package flightsHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rpparas/flight_log/database"
	"github.com/rpparas/flight_log/internals/model"
)

// GetFlights func gets all existing flights
// @Description Get all existing flights
// @Tags Flights
// @Accept json
// @Produce json
// @Success 200 {array} model.Flight
// @router /api/flights [get]
func GetFlights(c *fiber.Ctx) error {
	db := database.DB
	var flights []model.Flight

	// find all flights in the database
	db.Find(&flights)

	// If no flights is present return an error
	if len(flights) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No flights present", "data": nil})
	}

	// Else return flights
	return c.JSON(fiber.Map{"status": "success", "message": "Flights Found", "data": flights})
}

// CreateFlight func create a flight
// @Description Create a flight
// @Tags Flights
// @Accept json
// @Produce json
// @Success 200 {object} model.Flight
// @router /api/flights [post]
func CreateFlight(c *fiber.Ctx) error {
	db := database.DB
	flights := new(model.Flight)

	// Store the body in the flights and return error if encountered
	err := c.BodyParser(flights)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the flights
	flights.ID = uuid.New()
	// Create the Flight and return error if encountered
	err = db.Create(&flights).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create flights", "data": err})
	}

	// Return the created flights
	return c.JSON(fiber.Map{"status": "success", "message": "Created Flight", "data": flights})
}

// GetFlight func one flight by ID
// @Description Get one flights by ID
// @Tags Flight
// @Accept json
// @Produce json
// @Success 200 {object} model.Flight
// @router /api/flights/{id} [get]
func GetFlight(c *fiber.Ctx) error {
	db := database.DB
	var flights model.Flight

	// Read the param flightsId
	id := c.Params("flightsId")

	// Find the flights with the given Id
	db.Find(&flights, "id = ?", id)

	// If no such flights present return an error
	if flights.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No flights present", "data": nil})
	}

	// Return the flights with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Flights Found", "data": flights})
}
