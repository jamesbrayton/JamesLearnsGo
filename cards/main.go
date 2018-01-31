package main

import (
	"fmt"
)

func main() {
	// This is long hand for the same variable creation and assignment used below
	// var card string = "Ace of Spades"

	// := is only needed during creation of a new variable.  = is used for re-assignment or update of value
	card := "Ace of Spades"
	card = "Five of Diamonds"

	fmt.Println(card)
}
