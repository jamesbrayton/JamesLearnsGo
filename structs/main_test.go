package main

import (
	"testing"
)

// Test the update first name function
func TestUpdateFirstName(t *testing.T) {

	// Create a person to test with.
	testPerson := person{
		firstName: "Tester",
		lastName:  "McTestface",
		contact: contactInfo{
			email:   "partyanimal@beer.com",
			zipCode: 94000,
		},
	}

	// Assert that the first name of the test person so we can later assert that it changed.
	if testPerson.firstName != "Tester" {
		t.Errorf("Expected %v, but got %v", "Tester", testPerson.firstName)
	}

	// Update the person first name.
	testPerson.updateFirstName("Testy")

	// Assert that the first name of the test person was updated.
	if testPerson.firstName != "Testy" {
		t.Errorf("Expected %v, but got %v", "Testy", testPerson.firstName)
	}
}

// Tests the update slice function and helps prove the pass by reference behavior
func TestUpdateSlice(t *testing.T) {
	// Quick example of pass by reference// Example slice
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	// Call the update function.
	updateSlice(mySlice)

	// Assert the before so that we can assert the after.
	if mySlice[0] != "Bye" {
		t.Errorf("Expected %v, but got %v", "Bye", mySlice[0])
	}
}
