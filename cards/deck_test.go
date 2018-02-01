package main

import "testing"

func TestNewdeck(t *testing.T) {
	d := newDeck()

	// Verify that we have 52 cards in our deck.
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	// Verify some ordering of the deck by testing the first and last cards.
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected card "Ace of Spades", but got %v", d[0])
	}
}

func TestShuffleDeck(t *testing.T) {

}
