package flightsHandler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rpparas/flight_log/database"
	"github.com/rpparas/flight_log/internals/model"
)

// GetFlight func one flight by ID
// @Description Get one flight by ID
// @Tags Flight
// @Accept json
// @Produce json
// @Success 200 {object} model.Flight
// @router /api/v1/flight/{id} [get]
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

// GetFlights func gets all existing flights
// @Description Get all existing flights
// @Tags Flights
// @Accept json
// @Produce json
// @Success 200 {array} model.Flight
// @router /api/v1/flights [get]
func GetFlights(c *fiber.Ctx) error {
	// TODOs:
	// 1. limit number of results
	// 2. do pagination (token for next set of results)
	// 3. handle having more query strings

	db := database.DB
	var flights []model.Flight

	generation := parseQueryGeneration(c)
	if generation == -1 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "`generation` is not a valid numeric value [1 to 26]", "data": nil})
	}

	dateFrom, err := parseQueryDateTime(c, "from")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid date `from` provided. See RFC3339 for valid format", "data": nil})
	}

	dateTo, err := parseQueryDateTime(c, "to")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid date `to` provided. See RFC3339 for valid format", "data": nil})
	}

	if !isCompatibleDateRange(dateFrom, dateTo) {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "`from` date doesn't come after `to` date", "data": nil})
	}

	maxDurationMins := parseQuerymaxDurationMins(c)
	if maxDurationMins == -1 {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "`maxDurationMins` should be an integer value [1 to 30]", "data": nil})
	}

	// use scopes to chain conditions based on query strings in URL
	if generation > 0 {
		db = db.Scopes(GenerationEquals(generation))
	}

	if !dateFrom.IsZero() {
		db = db.Scopes(StartingFrom(dateFrom))
	}

	if !dateTo.IsZero() {
		db = db.Scopes(EndingIn(dateTo))
	}

	if maxDurationMins > 0 {
		// we need to convert these to epocs
		db = db.Scopes(ShorterThan(maxDurationMins))
	}

	// even if no query string is declared, then search everything
	db.Debug().Find(&flights)

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

func isCompatibleDateRange(dateFrom time.Time, dateTo time.Time) bool {
	if dateFrom.IsZero() {
		return true
	}
	if dateTo.IsZero() {
		return true
	}
	return dateFrom.Before(dateTo)
}

func parseQuerymaxDurationMins(c *fiber.Ctx) int {
	// Parses generation number if supplied in the URL as a query parameter
	// returns -1 if invalid
	// returns -2 if ignore
	// returns valid generation [1 to 26] if valid

	maxDurationMins := c.Query("maxDurationMins")
	if len(maxDurationMins) == 0 {
		return -2
	}

	intVal, err := strconv.Atoi(maxDurationMins)
	if err != nil {
		return -1
	}

	// TODO: figure out how to implement this via router as regex rule
	if intVal < 1 || intVal > 30 {
		return -1
	}
	return intVal
}
