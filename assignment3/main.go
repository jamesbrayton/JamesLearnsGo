// Create a program that reads the contents of a text file then prints its contents
// to the terminal.

// The file to open should be provided as an argument to the program when it is
// executed at the terminal.  For example, 'go run main.go mytextfile.txt' should open
// up the mytextfile.txt file.

// To read in the arguments provided to a program, you can reference the
// variable 'os.Args', which is a slice of type string.

// To open a file, check out the document for the 'Open' function in the 'os'
// package - https://golang.org/pkg/os/#Open

// What interfaces does the 'File' type implement?

// If the 'File' type implements the 'Reader' interface, you might be able to reuse
// the io.Copy function!

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Get the second argument in the Args slice; this would be the filename passed in.
	filename := os.Args[1]

	// Read in the file.
	file, err := os.Open(filename)

	// Check for an error.
	if err != nil {
		fmt.Printf("Unable to open the file %v.  Error: %v", filename, err)
		os.Exit(1)
	}

	// Use copy to output the contents of the file to the terminal.
	io.Copy(os.Stdout, file)
}
