package main

import (
	"fmt"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// variable _ is used when we are saying we understand a variable exists here (in this case our iterator), but we aren't actually
	// going to use it for anything.  This keeps the compiler from complaining.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// convention in a reciever is to refer to the instance with a one or two variable as in
// func (d deck) print()
// modified as thisDeck to relate to the idea of 'this' being use in Go
func (thisDeck deck) print() {
	for i, card := range thisDeck {
		fmt.Println(i, card)
	}
}

// Returns 2 decks.  The first deck is the first set of values up to the hand size, the second is what remains in the deck.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
