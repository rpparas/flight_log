package flightsHandler

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
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
// @router /api/v1/flight [post]
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
// @Accept text/csv
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
	path, err := saveTempFile(c, file)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Unable to save received CSV: " + err.Error(), "data": nil})
	}

	records, err := readRowsFromCsv(path)
	if err != nil {
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "Unable to read CSV attachment: " + err.Error(), "data": nil})
	}

	// TODO: add validation for records parsed from CSV
	db := database.DB
	numFlightsSaved := 0
	allErrors := []string{}

	for i, record := range records {
		flight, rowErrors := constructFlightRecord(record, i)
		if len(rowErrors) > 0 {
			continue
		}

		// TODO: optimize this INSERT query by inserting in batches
		err = db.Debug().Create(&flight).Error
		if err != nil {
			errorMsg := "Row " + strconv.Itoa(i+2) + " was not imported to DB " + err.Error()
			rowErrors = append(rowErrors, errorMsg)
		}

		allErrors = append(allErrors, rowErrors...)
		if err != nil {
			continue
		}

		numFlightsSaved += 1
	}

	if numFlightsSaved > 0 {
		// TODO: show the UUIDs of the flights created & the corresponding row?
		return c.Status(201).JSON(fiber.Map{"status": "success", "message": fmt.Sprint("Created ", numFlightsSaved, " Flights"), "data": nil, "errors": allErrors})
	} else {
		// TODO: determine if we should really use error 422 or something else
		return c.Status(422).JSON(fiber.Map{"status": "error", "message": "No flights were saved to the database.", "data": nil, "errors": allErrors})
	}

}

func saveTempFile(c *fiber.Ctx, file *multipart.FileHeader) (string, error) {
	path := "./tmp"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	path += "/" + file.Filename
	err := c.SaveFile(file, path)
	if err != nil {
		return "", nil
	}
	return path, nil
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
		errorMsg := "Row " + strconv.Itoa(i+2) + " has invalid robot ID " + record[0]
		errors = append(errors, errorMsg)
	}

	startTime, err := time.Parse(time.RFC3339, record[1])
	if err != nil {
		// TODO: create structure of what error message would be
		errorMsg := "Row " + strconv.Itoa(i+2) + " has invalid start_time " + record[1]
		errors = append(errors, errorMsg)
	}

	endTime, err := time.Parse(time.RFC3339, record[2])
	if err != nil {
		// TODO: create structure of what error message would be
		errorMsg := "Row " + strconv.Itoa(i+2) + " has invalid end_time " + record[2]
		errors = append(errors, errorMsg)
	}

	lat, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		// TODO: create structure of what error message would be
		errorMsg := "Row " + strconv.Itoa(i+2) + " has invalid latitude " + record[3]
		errors = append(errors, errorMsg)
	}

	lng, err := strconv.ParseFloat(record[4], 64)
	if err != nil {
		// TODO: create structure of what error message would be
		errorMsg := "Row " + strconv.Itoa(i+2) + " has invalid longitude " + record[4]
		errors = append(errors, errorMsg)
	}

	if len(errors) > 0 {
		return model.Flight{}, errors
	}

	flight := model.Flight{
		RobotID:   robotId,
		StartTime: startTime,
		EndTime:   endTime,
		Lat:       lat,
		Lng:       lng,
	}

	return flight, errors
}
