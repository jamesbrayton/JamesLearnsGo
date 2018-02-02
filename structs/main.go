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
	fmt.Printf("%+v", alex)

	fmt.Printf("%+v", jim)
}
