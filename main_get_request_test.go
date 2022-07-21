package main

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/rpparas/flight_log/database"
	"github.com/stretchr/testify/assert"
)

func TestIndexRoute(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "index route",
			route:         "/",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "TODO: Add documentation how to use Drone Log API",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	app := Setup()

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		res, err := app.Test(req, -1)

		// verify that no error occured, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		// Read the response body
		body, err := ioutil.ReadAll(res.Body)

		// Reading the response body should work everytime, such that
		// the err variable should be nil
		assert.Nilf(t, err, test.description)

		// Verify, that the reponse body equals the expected body
		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}

func TestGetFlights(t *testing.T) {
	tests := []TestCase{
		{
			description:     "GET Flights with at least 1 result",
			route:           "/api/v1/flights",
			expectedError:   false,
			expectedCode:    200,
			expectedMessage: "Flights Found",
		},
		{
			description:     "GET Flights for robot generation without any result",
			route:           "/api/v1/flights?generation=26",
			expectedError:   false,
			expectedCode:    200,
			expectedMessage: "No flights present",
		},
		{
			description:     "GET Flights for with invalid generation",
			route:           "/api/v1/flights?generation=0",
			expectedError:   false,
			expectedCode:    400,
			expectedMessage: "`generation` is not a valid numeric value [1 to 26]",
		},
		{
			description:     "GET Flights for with invalid generation",
			route:           "/api/v1/flights?generation=99999",
			expectedError:   false,
			expectedCode:    400,
			expectedMessage: "`generation` is not a valid numeric value [1 to 26]",
		},
		{
			description:     "GET Flights for with invalid generation",
			route:           "/api/v1/flights?generation=not-a-valid-value",
			expectedError:   false,
			expectedCode:    400,
			expectedMessage: "`generation` is not a valid numeric value [1 to 26]",
		},
		{
			description:     "GET Flights with valid `from` date",
			route:           "/api/v1/flights?from=2018-01-01T00:00:00Z",
			expectedError:   false,
			expectedCode:    200,
			expectedMessage: "Flights Found",
		},
		{
			description:     "GET Flights with invalid `from` date",
			route:           "/api/v1/flights?from=Apr 1 2022",
			expectedError:   false,
			expectedCode:    400,
			expectedMessage: "Invalid date `from` provided. See RFC3339 for valid format",
		},
		{
			description:     "GET Flights with generation, `from` date & to `to` date",
			route:           "/api/v1/flights?generation=1&from=2018-01-01T00:00:00Z&to=2023-07-03T00:00:00Z",
			expectedError:   false,
			expectedCode:    200,
			expectedMessage: "Flights Found",
		},
		{
			description:     "GET Flights with incompatible `from` & `to` dates",
			route:           "/api/v1/flights?to=2022-07-01T00:00:00Z&from=2023-07-03T00:00:00Z",
			expectedError:   false,
			expectedCode:    400,
			expectedMessage: "`from` date doesn't come after `to` date",
		},
		{
			description:     "GET Flights with maxDurationMins",
			route:           "/api/v1/flights?maxDurationMins=16",
			expectedError:   false,
			expectedCode:    200,
			expectedMessage: "Flights Found",
		},
	}

	executeGetTests(t, tests)
}

func executeGetTests(t *testing.T, tests []TestCase) {
	database.ConnectDB()

	// Setup the app as it is done in the main function
	app := Setup()

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		res := expectedMatchesActual(t, test, app, req)
		if !res {
			continue
		}
	}
}
