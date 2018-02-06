package main

import "testing"

// Test function for testing the getArea() function against a triangle struct.
func TestGetAreaOfATriangle(t *testing.T) {
	// Set up our base and height values to test with.
	testHeight := float64(4)
	testBase := float64(3)
	calculatedArea := 0.5 * testHeight * testBase

	// Create a test struct to work with.
	testTriangle := triangle{height: testHeight, base: testBase}

	// Test the getArea() function
	testedArea := testTriangle.getArea()

	// Assert the result
	if calculatedArea != testedArea {
		t.Errorf("Error.  Expected %v, but got %v.", calculatedArea, testedArea)
	}
}

// Test function for testing the getArea() function against a square struct.
func TestGetAreaOfASquare(t *testing.T) {
	// Set up our base and height values to test with.
	testSideLength := float64(4)
	calculatedArea := testSideLength * testSideLength

	// Create a test struct to work with.
	testSquare := square{sideLength: testSideLength}

	// Test the getArea() function
	testedArea := testSquare.getArea()

	// Assert the result
	if calculatedArea != testedArea {
		t.Errorf("Error.  Expected %v, but got %v.", calculatedArea, testedArea)
	}
}
