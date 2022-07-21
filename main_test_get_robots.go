package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRobots(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/api/v1/roboaat",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/invalid-url",
			expectedCode: 404,
		},
	}

	app := Setup()

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)

		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
