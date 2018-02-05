package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex37@lougle.com"
	alex.contact.zipCode = 12345

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "partyanimal@beer.com",
			zipCode: 94000,
		},
	}

	fmt.Println(alex)

	// Prints out all the field names plus their values.
	fmt.Printf("%+v\n", alex)

	// Use our new print function.
	jim.print()

	// NOTES:
	// varaiable - Give me the value this variable has at its memory address. (value)
	// &variable - Give me the memory address of the value this variable is pointing at (reference).
	// *pointer - Give me the value this memory address is pointing at (pointer).
	// In other words, use *address to turn an address into a value
	// use &value to turn a value into an address

	// Update jim using the pointer.
	jim.updateFirstName("Jimbo")

	jim.print()

	// Quick example of pass by reference// Example slice
	mySlice := []string{"Hi", "There", "How", "Are", "You"}

	// Call the update function.
	updateSlice(mySlice)

	// Output the result
	fmt.Println(mySlice)
}

// Update name on the person using pointers.
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Update function
func updateSlice(s []string) {
	s[0] = "Bye"
}

// Print the person object formatted to include all possible value names.
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
