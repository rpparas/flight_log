package flightsHandler

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"

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
	flight := new(model.Flight)

	// Store the body in the flights and return error if encountered
	err := c.BodyParser(flight)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": err.Error() + " Review your input", "data": err})
	}

	// TODO: add error-handling for missing or invalid data in payload
	flight.ID = uuid.New()

	err = db.Create(&flight).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create flight", "data": err})
	}

	// Return the created flight
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created Flight", "data": flight})
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
	err = c.SaveFile(file, path)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Unable to save received CSV", "data": nil})
	}

	records, err := readRowsFromCsv(path)
	if err != nil {
		log.Fatal(err)
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Unable to read CSV attachment", "data": nil})
	}

	// TODO: add validation for records parsed from CSV
	db := database.DB
	numFlightsSaved := 0

	for i, record := range records {
		flight, errors := constructFlightRecord(record, i)
		if len(errors) > 0 {
			log.Println(errors)
			continue
		}

		err = db.Create(&flight).Error
		if err != nil {
			log.Println(err)
			// TODO: skip ?
			continue
			// return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create flight", "data": err})
		}

		numFlightsSaved += 1
	}

	// TODO: save to db
	// TODO: Return the created flights
	// TODO: show errors
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created Flights", "data": nil, "errors": nil})
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

func constructFlightRecord(record []string, i int) (model.Flight, []string) {
	robotId, err := uuid.Parse(record[0])
	errors := []string{}

	if err != nil {
		// TODO: create structure of what error message would be
		errorMsg := "Row " + strconv.Itoa(i+2) + " invalid robot ID " + record[0]
		errors = append(errors, errorMsg)
	}

	startTime, err := time.Parse(time.RFC3339, record[1])
	if err != nil {
		// TODO: create structure of what error message would be
		errorMsg := "Row " + strconv.Itoa(i+2) + " invalid start_time " + record[1]
		errors = append(errors, errorMsg)
	}

	endTime, err := time.Parse(time.RFC3339, record[2])
	if err != nil {
		// TODO: create structure of what error message would be
		errorMsg := "Row " + strconv.Itoa(i+2) + " invalid end_time " + record[2]
		errors = append(errors, errorMsg)
	}

	lat, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		// TODO: create structure of what error message would be
		errorMsg := "Row " + strconv.Itoa(i+2) + " invalid latitude " + record[3]
		errors = append(errors, errorMsg)
	}

	lng, err := strconv.ParseFloat(record[4], 64)
	if err != nil {
		// TODO: create structure of what error message would be
		errorMsg := "Row " + strconv.Itoa(i+2) + " invalid longitude " + record[4]
		errors = append(errors, errorMsg)
	}

	if len(errors) > 0 {
		return model.Flight{}, errors
	}
	// TODO: construct record to save to DB
	flight := model.Flight{
		RobotID:   robotId,
		StartTime: startTime,
		EndTime:   endTime,
		Lat:       lat,
		Lng:       lng,
	}

	return flight, errors
}
