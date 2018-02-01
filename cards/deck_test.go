package main

import "testing"

func TestNewdeck(t *testing.T) {
	d := newDeck()

	// Verify that we have 52 cards in our deck.
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
}
