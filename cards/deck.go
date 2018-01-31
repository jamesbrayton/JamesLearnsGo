package main

import (
	"fmt"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

// convention in a reciever is to refer to the instance with a one or two variable as in
// func (d deck) print()
// modified as thisDeck to relate to the idea of 'this' being use in Go
func (thisDeck deck) print() {
	for i, card := range thisDeck {
		fmt.Println(i, card)
	}
}
