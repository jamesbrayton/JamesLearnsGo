// In this assignment you'll get some practice with slices and for loops.

// Here are the directions:

// Create a new project folder to house this new project and create a 'main.go' file inside of it.
// In the main.go file, define a function called 'main'.  Remember that the 'main' function will be called automatically.
// In the main function, create a slice of ints from 0 through 10
// Iterate through the slice with a for loop.  For each element, check to see if the number is even or odd.
// If the value is even, print out "even".  If it is odd, print out "odd"
// Run your code from the terminal by changing into your project directory then running 'go run main.go'.
// Your output should look like this:

// 0 is even
// 1 is odd
// 2 is even
// 3 is odd
// 4 is even
// 5 is odd
// 6 is even
// 7 is odd
// 8 is even
// 9 is odd
// 10 is even

package main

import (
	"fmt"
)

// Main function
func main() {
	// slice of ints containing 1 though 10
	ints := generateSlice(10)

	// for loop iteration of the slice.
	for i := range ints {
		// check to see if the value is even or odd
		if ints[i]%2 == 0 {
			fmt.Println(ints[i], "is even")
		} else {
			fmt.Println(ints[i], "is odd")
		}
	}
}

// Function for generating a slice of ints
func generateSlice(numberOfValuesInSlice int) []int {
	slice := make([]int, numberOfValuesInSlice)

	for i := range slice {
		slice[i] = i + 1
	}

	return slice
}
