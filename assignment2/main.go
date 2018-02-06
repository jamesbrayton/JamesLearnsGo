package main

import (
	"fmt"
)

// Write a program that creates two custom struct types called
// 'triangle' and 'square'.

// The 'square' type should be a struct with a field called 'sideLength'
// of type float64.

// the 'triangle' type should be a struct with a field called 'height' of
// type float64 and a field of type 'base' of type float64.

// Both types should have a function called 'getArea' that returns the
// calculated area of the square or triangle.

// Area of a triangle = 0.5 * base * height.
// Area of a square = sideLength * sideLength

// Add a 'shape' interface that defines a function called 'printArea'.
// This function should calculate the area of the given shape and
// print it out to the terminal. Design the interface so that the 'printArea'
// function can be called with either a triangle or a square.

// Square struct
type square struct {
	sideLength float64
}

// Triangle struct
type triangle struct {
	height float64
	base   float64
}

// Shape interface
type shape interface {
	getArea() float64
}

// Main function
func main() {
	// Demonstrate the square
	sq := square{sideLength: 4}

	// Print out the area of the square.
	fmt.Println("The area of a square with sides of length 4 is:")
	printArea(sq)

	// Demonstrate the triangle
	tr := triangle{base: 4, height: 3}

	// Print out the area of the triangle.
	fmt.Println("The area of a triangle with base of 4 and height of 3 is:")
	printArea(tr)
}

// Function to return the area of a triangle.
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

// Function to return the area of a square.
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

// Interface function to write out the area of a shape to the commandline.
func printArea(s shape) {
	fmt.Println(s.getArea())
}
