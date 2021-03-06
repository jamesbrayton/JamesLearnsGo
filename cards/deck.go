package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
			cards = append(cards, value+" of "+suit)
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

// Returns the deck as a single comma delimitted string.
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Returns the deck as a byte slice.
func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

// Attempts to write the deck to the file system.
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, d.toByteSlice(), 0666)
}

// Attempts to read a deck from a file.  If it is unable to do so,
// writes an error to the terminal and exits with an error code of 1.
func newDeckFromFile(filename string) deck {
	deckAsByteSlice, err := ioutil.ReadFile(filename)

	// If we encounter and error trying to open the file, write the
	// error to the terminal and exit.
	if err != nil {
		fmt.Println("ERROR: ", err)
		os.Exit(1)
	}

	// turn the byte slice into a deck.
	return deck(strings.Split(string(deckAsByteSlice), ","))
}

// Shuffles a deck of cards.
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	for i := range d {
		newPosition := random.Intn(len(d) - 1)

		// One line swap.
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
