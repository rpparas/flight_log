package main

import (
	"net/http"
	"strings"
	"testing"

	"github.com/rpparas/flight_log/database"
)

func TestPostFlightsJson(t *testing.T) {
	tests := []TestCase{
		{
			description: "POST single Flight",
			route:       "/api/v1/flights",
			payload: `{
				"robotId": "e570b6c0-9bb0-47c9-a358-b984ed402406",
				"startTime": "2022-07-18T01:00:01+00:00"
			}`,
			expectedError:   false,
			expectedCode:    201,
			expectedMessage: "Created Flight",
		},
		{
			description: "POST duplicate Flight",
			route:       "/api/v1/flights",
			payload: `{
				"robotId": "e570b6c0-9bb0-47c9-a358-b984ed402406",
				"startTime": "2022-07-18T01:00:01+00:00"
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
		compareTestResults(t, test, app, req)
	}

}
