package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type Result struct {
	Message string
}

type TestCase struct {
	description string
	route       string
	payload     string
	filepath    string

	// Expected output
	expectedError   bool
	expectedCode    int
	expectedMessage string
}

func compareTestResults(t *testing.T, test TestCase, app *fiber.App, req *http.Request) {
	// verify that no error occured, that is not expected
	res, err := app.Test(req, -1)
	assert.Equalf(t, test.expectedError, err != nil, test.description)

	// As expected errors lead to broken responses, the next
	// test case needs to be processed
	if test.expectedError {
		return
	}

	// Verify if the status code is as expected
	assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	// parse api result as JSON
	var apiResult Result
	err = json.Unmarshal(body, &apiResult)
	if err != nil {
		fmt.Println(err)
	}

	assert.Nilf(t, err, test.description)

	assert.Equalf(t, test.expectedMessage, apiResult.Message, test.description)

}
