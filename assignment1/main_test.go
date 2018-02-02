package main

import "testing"

// Test the generate slice function
func TestGenerateSlice(t *testing.T) {
	// number of values in slice to generate.
	sizeOfGeneratedSlice := 5

	// generate a slice to test with.
	testSlice := generateSlice(sizeOfGeneratedSlice)

	// Assert that the slice is the correct size.
	if len(testSlice) != sizeOfGeneratedSlice {
		t.Errorf("Expected size of slice should be %v, but is %v", sizeOfGeneratedSlice, testSlice)
	}

	// Assert that the slice has the correct values.
	for i := range testSlice {
		if testSlice[i] != i+1 {
			t.Errorf("Expected %v at index %v, but got %v", i+1, i, testSlice[i])
		}
	}
}
