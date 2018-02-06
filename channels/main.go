package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Interface for defining an http client.  This is used for testing later.
type httpClient interface {
	Get(url string) (*http.Response, error)
}

// Struct for holding checkLink response data.
type checkResponse struct {
	success bool
	message string
}

// Main function.
func main() {
	links := getLinks()

	// Create a waitgroup to ensure all go routines finish before exiting the program.
	var wg sync.WaitGroup

	// For loop to check all the links.
	for _, link := range links {
		// Increment the waitgroup counter.
		wg.Add(1)

		// Launch a goroutine to check the link and write the response.
		go func(link string) {

			// Decerment the counter when the goroutine completes.
			defer wg.Done()

			// Make the actual call to check link and print the response.
			fmt.Println(checkLink(new(http.Client), link))
		}(link)
	}

	// Wait for all the goroutines to complete.
	wg.Wait()
}

// Check link function that returns a boolean value indicating the success of the call
// and a string message explaining the result.  This takes an httpClient interface and
// a string for the link to call.
func checkLink(hc httpClient, link string) checkResponse {
	// Make a get request against the link.
	_, err := hc.Get(link)

	// If we encountered an error, return false and an error message.
	if err != nil {
		return checkResponse{success: false, message: link + " might be down."}
	}

	// If no error, return true and a message.
	return checkResponse{success: true, message: link + " is up."}
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
