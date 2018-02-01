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

func TestShuffleDeck(t *testing.T) {

}
