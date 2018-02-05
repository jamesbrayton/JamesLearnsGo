package main

import (
	"fmt"
)

func main() {
	// Literal syntax for creating a map.
	colors1 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// Print out the contents of the colors1 map.
	fmt.Println(colors1)

	// Alternate syntax
	var colors2 map[string]string

	// Print out the contents of the colors2 map.
	fmt.Println(colors2)

	// Third syntax using make()
	colors3 := make(map[string]string)

	// Print out the contents of the colors3 map.
	fmt.Println(colors3)

	// Add "white" to the colors3 map.
	colors3["white"] = "#ffffff"

	// print out the colors3 map to see that white was added.
	fmt.Println(colors3)

	// Delete function example
	delete(colors3, "white")

	// print out the colors3 map to see that white was deleted.
	fmt.Println(colors3)
}
