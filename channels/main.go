package main

import (
	"fmt"
	"net/http"
	"time"
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

	// Make a channel to allow inter-goroutine communication
	c := make(chan string)

	// Make the httpClient to use.
	hc := new(http.Client)

	// For loop to check all the links.
	for _, link := range links {
		// Launch a goroutine to check the link and write the response.
		go checkLink(hc, link, c)
	}

	// Setup the channel to output messages to the channel to the terminal.
	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			fmt.Println(checkLink(hc, l, c))
		}(l)
	}
}

// Check link function that returns a boolean value indicating the success of the call
// and a string message explaining the result.  This takes an httpClient interface and
// a string for the link to call.
func checkLink(hc httpClient, link string, c chan string) checkResponse {
	// Make a get request against the link.
	_, err := hc.Get(link)

	// Response object to store our success of failure in.
	var resp checkResponse

	// If we encountered an error, return false and an error message.
	if err != nil {
		resp = checkResponse{success: false, message: link + " might be down."}
	} else {
		// If no error, return true and a message.
		resp = checkResponse{success: true, message: link + " is up."}
	}

	// Send the link back into the channel to check again.
	c <- link

	// Return the response.
	return resp
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
