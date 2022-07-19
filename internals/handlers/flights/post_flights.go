package flightsHandler

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rpparas/flight_log/database"
	"github.com/rpparas/flight_log/internals/model"
)

// CreateFlight func create a flight
// @Description Create a flight
// @Tags Flights
// @Accept json
// @Produce json
// @Success 201 {object} model.Flight
// @router /api/v1/flights [post]
func CreateFlight(c *fiber.Ctx) error {
	db := database.DB
	flights := new(model.Flight)

	// Store the body in the flights and return error if encountered
	err := c.BodyParser(flights)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": err.Error() + " Review your input", "data": err})
	}

	// TODO: add error handling for missing or invalid data in payload
	flights.ID = uuid.New()

	err = db.Create(&flights).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create flight", "data": err})
	}

	// Return the created flight
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created Flight", "data": flights})
}

// CreateFlight func create a flight
// @Description Create a flight
// @Tags Flights
// @Accept json
// @Produce json
// @Success 201 {object} model.Flight
// @router /api/v1/flights/csv [post]
func CreateFlights(c *fiber.Ctx) error {
	// Inserts flights to flight log
	// Assumptions:
	// 1. Correctly-formatted CSV
	// 2. Complete meta data (robot ID, start_time, end_time, lat, lng)
	// 3. Robot ID already exists in DB
	// 4.
	c.Accepts("text/csv")
	file, err := c.FormFile("document")
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Missing CSV attachment", "data": nil})
	}

	path := "./tmp/" + file.Filename
	log.Println(file.Filename)
	c.SaveFile(file, path)
	records, err := readRowsFromCsv(path)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Unable to read CSV attachment", "data": nil})
	}

	errors := []string{}

	for i, record := range records {
		robotId, err := uuid.Parse(record[0])
		if err != nil {
			// TODO: create structure of what error message would be
			errorMsg := "Skipped row " + strconv.Itoa(i+2) + " invalid robot ID " + record[0]
			errors = append(errors, errorMsg)
			continue
		}
		// TODO: construct record to save to DB
		flight := model.Flight{
			RobotID: robotId,
		}
		log.Println(flight)
	}

	// TODO: save to db
	// db := database.DB

	// Return the created flights
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created Flights", "data": nil, "errors": errors})
}

func readRowsFromCsv(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
