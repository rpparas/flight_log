package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type Result struct {
	Message string
	Errors  []string
}

type TestCase struct {
	description string
	route       string
	payload     string
	filepath    string

	// Expected output
	expectedError   bool
	expectedErrors  []string
	expectedCode    int
	expectedMessage string
}

func expectedMatchesActual(t *testing.T, test TestCase, app *fiber.App, req *http.Request) bool {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// verify that no error occured, that is not expected
	res, err := app.Test(req, -1)
	assert.Equalf(t, test.expectedError, err != nil, test.description)

	if test.expectedError {
		return false
	}

	// Verify if the status code is as expected
	assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	// parse api result as JSON
	var apiResult Result
	err = json.Unmarshal(body, &apiResult)
	if err != nil {
		log.Println(err)
	}

	assert.Nilf(t, err, test.description)

	assert.Equalf(t, test.expectedMessage, apiResult.Message, test.description)
	// TODO: assert errors if they exist
	// assert.Equalf(t, test.expectedErrors, apiResult.Errors, test.description)

	return true
}
