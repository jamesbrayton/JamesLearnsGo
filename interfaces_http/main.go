package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}

func main() {
	// Make a get request against google.com
	resp, err := http.Get("http://www.google.com")

	// Check for an error.
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Make an empty byte slice with room for 99999 elements.
	body := make([]byte, 9999)

	// Read the body of the response into our body byte slice.
	numOfBytesRead, err := resp.Body.Read(body)

	// Check for errors.
	if (numOfBytesRead == 0) || (err != nil) {
		fmt.Println("Number of bytes read: ", numOfBytesRead)
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Print out the response.
	// fmt.Println(string(body))

	// One line function call to do the samething as lines 20 - 34 above.
	// io.Copy(os.Stdout, resp.Body)

	// Create an instance of our custom writer.
	cw := customWriter{}

	// Call copy with our custom writer.
	io.Copy(cw, resp.Body)
}

func (customWriter) Write(bs []byte) (int, error) {

	// Write out the contents of the byte slice to the command line.
	fmt.Println(string(bs))

	// Write out the number of bytes written to the the terminal.
	fmt.Println("Just wrote", len(bs), "bytes.")

	// Return the length of our byteslice that was read out, and nil for no errors.
	return len(bs), nil
}
