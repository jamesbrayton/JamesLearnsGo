package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A mock http object for testing.
type MockHttpClient struct {
}

func (m MockHttpClient) Get(url string) (resp *http.Response, err error) {
	return new(http.Response), nil
}

// Tests the check link function.  The function returns a boolean value
// indicating the success of the call and a string message explaining the result.
func TestCheckLink(t *testing.T) {
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
