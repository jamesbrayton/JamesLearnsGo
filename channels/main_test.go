package main

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A mock http object for testing.
type MockHttpClient struct {
}

// Function for overriding Get behavior on an http.client in order to test.
func (m MockHttpClient) Get(url string) (resp *http.Response, err error) {
	// Create failure scenario to test both positive and negative behavior.
	if url == "https://should.fail" {
		return new(http.Response), errors.New("no mi gusta")
	}

	return new(http.Response), nil
}

// Tests the check link function.  The function returns a boolean value
// indicating the success of the call and a string message explaining the result.
func TestCheckLinkSucceeds(t *testing.T) {
	// Mock http client.
	mockClient := MockHttpClient{}

	// Link to test with the mock client.
	testLink := "https://yourMom.com"

	// Call the code we are testing.
	success, message := checkLink(mockClient, testLink)

	// Assert the expectations we met.
	assert.True(t, success)
	assert.Contains(t, message, testLink)
}

// Tests the check link function under a failure scenario.  The function returns a
// boolean value indicating the success of the call and a string message explaining
// the result.
func TestCheckLinkFail(t *testing.T) {
	// Mock http client.
	mockClient := MockHttpClient{}

	// Link to test with the mock client.
	testLink := "https://should.fail"

	// Call the code we are testing.
	success, message := checkLink(mockClient, testLink)

	// Assert the expectations we met.
	assert.False(t, success)
	assert.Contains(t, message, testLink)
}

// Test the getLinks() function.
func TestGetLinks(t *testing.T) {
	expectedLinks := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	actualLinks := getLinks()

	assert.Equal(t, expectedLinks, actualLinks)
}
