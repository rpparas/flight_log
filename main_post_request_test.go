package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/rpparas/flight_log/database"
)

func TestPostFlightsJson(t *testing.T) {
	fakeTime := time.Now().Format(time.RFC3339)
	tests := []TestCase{
		{
			description: "POST single Flight",
			route:       "/api/v1/flight",
			payload: `{
				"robotId": "e570b6c0-9bb0-47c9-a358-b984ed402406",
				"startTime": "` + fakeTime + `"
			}`,
			expectedError:   false,
			expectedCode:    201,
			expectedMessage: "Created Flight",
		},
		{
			description: "POST duplicate Flight",
			route:       "/api/v1/flight",
			payload: `{
				"robotId": "e570b6c0-9bb0-47c9-a358-b984ed402406",
				"startTime": "` + fakeTime + `"
			}`,
			expectedError:   false,
			expectedCode:    500,
			expectedMessage: "Could not create flight",
		},
	}
	executePostTestsJson(t, tests)
}

func executePostTestsJson(t *testing.T, tests []TestCase) {
	database.ConnectDB()

	// Setup the app as it is done in the main function
	app := Setup()

	for _, test := range tests {
		req, _ := http.NewRequest(
			"POST",
			test.route,
			strings.NewReader(test.payload),
		)

		req.Header.Add("Content-Type", "application/json")
		res := expectedMatchesActual(t, test, app, req)
		if !res {
			continue
		}
	}

}

func TestPostFlightsCsv(t *testing.T) {
	twoErrors := []string{"error placeholder1", "error placeholder2"}
	tests := []TestCase{
		{
			description:     "POST bulk Flights via CSV",
			route:           "/api/v1/flights/csv",
			filepath:        "./examples/flights.csv",
			expectedError:   false,
			expectedCode:    201,
			expectedMessage: "Created 2 Flights",
			expectedErrors:  []string{},
		},
		{
			description:     "POST bulk duplicate Flights via CSV",
			route:           "/api/v1/flights/csv",
			filepath:        "./examples/flights.csv",
			expectedError:   false,
			expectedCode:    400,
			expectedMessage: "No flights were saved to the database.",
			expectedErrors:  twoErrors,
		},
		{
			description:     "POST empty file (no contents)",
			route:           "/api/v1/flights/csv",
			filepath:        "./examples/invalid.csv",
			expectedError:   false,
			expectedCode:    400,
			expectedMessage: "Unable to read CSV attachment: EOF",
			expectedErrors:  []string{},
		},
	}
	executePostTestsCsv(t, tests)
}

func executePostTestsCsv(t *testing.T, tests []TestCase) {
	database.ConnectDB()

	// Setup the app as it is done in the main function
	app := Setup()

	for _, test := range tests {
		payload := &bytes.Buffer{}
		writer := multipart.NewWriter(payload)
		file, errFile := os.Open(test.filepath)
		if errFile != nil {
			return
		}
		defer file.Close()

		part1, _ := writer.CreateFormFile("document", filepath.Base(test.filepath))
		_, errFile = io.Copy(part1, file)
		if errFile != nil {
			log.Println(errFile)
			return
		}
		err := writer.Close()
		if err != nil {
			log.Println(err)
			return
		}

		req, _ := http.NewRequest(
			"POST",
			test.route,
			payload,
		)

		req.Header.Add("Content-Type", writer.FormDataContentType())
		res := expectedMatchesActual(t, test, app, req)
		if !res {
			continue
		}
	}

}
