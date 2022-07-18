package flightsHandler

import (
	"strconv"
	"time"

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
// @router /api/v1/flights [get]
func GetFlights(c *fiber.Ctx) error {
	db := database.DB
	var flights []model.Flight

	generation := parseQueryGeneration(c)
	dateFrom, err := parseQueryDateTime(c, "from")
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Invalid date `from` provided. See RFC3339 for valid format", "data": nil})
	}

	if generation == -1 {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "`generation` is not a valid numeric value [1 to 26]", "data": nil})
	}

	// TODO: figure out how to chain these queries
	if generation > 0 {
		db.Model(&model.Flight{}).Joins("left join robots on flights.robot_id = robots.id").Where("generation = ?", generation).Scan(&flights)
	} else if !dateFrom.IsZero() {
		db.Model(&model.Flight{}).Where("start_time >= ?", dateFrom).Scan(&flights)
	} else {
		// if no query string is declared, then search everything
		db.Find(&flights)
	}

	// If no flights is present return no content
	if len(flights) == 0 {
		// TODO: Determine appropriate error code, 204 or 404 instead of 200
		return c.Status(200).JSON(fiber.Map{"status": "no content", "message": "No flights present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Flights Found", "data": flights})
}

func parseQueryGeneration(c *fiber.Ctx) int {
	// Parses generation number if supplied in the URL as a query parameter
	// returns -1 if invalid
	// returns -2 if ignore
	// returns valid generation [1 to 26] if valid

	generation := c.Query("generation")
	if len(generation) == 0 {
		return -2
	}

	intVal, err := strconv.Atoi(generation)
	if err != nil {
		return -1
	}

	// TODO: figure out how to implement this via router as regex rule
	if intVal < 1 || intVal > 26 {
		return -1
	}
	return intVal
}

func parseQueryDateTime(c *fiber.Ctx, queryName string) (_ time.Time, err error) {
	// Parses datetime string from `queryName` URL param based off RFC3339 format
	// returns "time-zero" equivalent if invalid
	// returns valid Time if successfully parsed

	dateTimeStr := c.Query(queryName)

	if len(dateTimeStr) == 0 {
		return
	}

	dateObj, err := time.Parse(time.RFC3339, dateTimeStr)
	if err != nil {
		return
	}
	return dateObj, nil
}

// CreateFlight func create a flight
// @Description Create a flight
// @Tags Flights
// @Accept json
// @Produce json
// @Success 200 {object} model.Flight
// @router /api/v1/flights [post]
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
// @Description Get one flight by ID
// @Tags Flight
// @Accept json
// @Produce json
// @Success 200 {object} model.Flight
// @router /api/v1/flights/{id} [get]
func GetFlight(c *fiber.Ctx) error {
	db := database.DB
	var flights model.Flight

	// Read the param flightsId
	id := c.Params("flightsId")

	// Find the flights with the given Id
	db.Find(&flights, "id = ?", id)

	// If no such flights present return an error
	if flights.ID == uuid.Nil {
		// TODO: Determine appropriate error code, 204 or 404 instead of 200
		return c.Status(200).JSON(fiber.Map{"status": "no content", "message": "No flight found", "data": nil})
	}

	// Return the flights with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Flights Found", "data": flights})
}
