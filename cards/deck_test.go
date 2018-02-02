package main

import "testing"

// Verify that we have 52 cards in our deck.
func TestNewDeckHasCorrectNumberOfCards(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}

// Verify some strict ordering of the deck by testing the first card.
func TestFirstCardInNewDeckIsAceOfSpades(t *testing.T) {
	d := newDeck()

	testCard := "Ace of Spades"

	if d[0] != testCard {
		t.Errorf("Expected card %v, but got %v", testCard, d[0])
	}
}

// Verify some strict ordering of the deck by testing the last card.
func TestLastCardInNewDeckIsKingOfClubs(t *testing.T) {
	d := newDeck()

	testCard := "King of Clubs"

	if d[len(d)-1] != testCard {
		t.Errorf("Expected card '%v', but got %v", testCard, d[len(d)-1])
	}
}

// Test that shuffle shuffles a deck.
func TestShuffleDeck(t *testing.T) {
	// Create 2 new decks
	d1, d2 := newDeck(), newDeck()

	// Assert the ordering of d1 and d2 are the same.
	for i := range d1 {
		if d1[i] != d2[i] {
			t.Errorf("Expected card '%v', but got %v", d1[i], d2[i])
		}
	}

	// Shuffle the first deck.

}
