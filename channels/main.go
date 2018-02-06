package main

import (
	"fmt"
	"net/http"
)

// Interface for defining an http client.  This is used for testing later.
type httpClient interface {
	Get(url string) (*http.Response, error)
}

// Main function.
func main() {
	links := getLinks()

	for _, link := range links {
		fmt.Println(checkLink(new(http.Client), link))
	}
}

// Check link function that returns a boolean value indicating the success of the call
// and a string message explaining the result.  This takes an httpClient interface and
// a string for the link to call.
func checkLink(hc httpClient, link string) (bool, string) {
	// Make a get request against the link.
	_, err := hc.Get(link)

	// If we encountered an error, return false and an error message.
	if err != nil {
		return false, link + " might be down."
	}

	// If no error, return true and a message.
	return true, link + " is up."
}

// Function for returning links.
func getLinks() []string {
	return []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}
}