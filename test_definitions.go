package main

type Result struct {
	Message string
}

type TestCase struct {
	description string
	route       string
	payload     string

	// Expected output
	expectedError   bool
	expectedCode    int
	expectedMessage string
}
