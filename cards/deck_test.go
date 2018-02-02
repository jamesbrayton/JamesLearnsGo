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

// Test the deal function on a deck of cards.
func TestDeal(t *testing.T) {
	// Get a new deck to test with.
	testDeck := newDeck()

	// Pick 7 cards to be the handsize.
	handSize := 7

	// Deal a new hand from the test deck.
	testHand, testDeck := deal(testDeck, handSize)

	// Assert that the test hand is the correct hand size.
	if len(testHand) != handSize {
		t.Errorf("Expected '%v' number of cards in the hand, but got %v", handSize, len(testHand))
	}

	// Verify that the testdeck contains the correct number of cards remaining.
	if len(testDeck) != (52 - handSize) {
		t.Errorf("Expected '%v' number of cards left in the deck, but got %v", (52 - handSize), len(testDeck))
	}
}

// Test the to string function of a deck
func TestToString(t *testing.T) {
	// Create a test deck to test with.
	// Overwrite the value in the test deck to only two cards for easier testing.
	// (I'm too lazt right now to type out all 52 cards for testing.)
	testDeck := deck{"Ace of Spades", "Two of Diamonds"}

	testStringShouldBe := "Ace of Spades,Two of Diamonds"

	// Call toString() on the deck.
	testDeckAsString := testDeck.toString()

	if testDeckAsString != testStringShouldBe {
		t.Errorf("Expected '%v', but got %v", testStringShouldBe, testDeckAsString)
	}
}

// Test that shuffle shuffles a deck.
func TestShuffleDeck(t *testing.T) {
	// Create two new decks
	testDeck, unshuffledTestDeck := newDeck(), newDeck()

	// Shuffle the first deck.
	testDeck.shuffle()

	// Assert at least one card in the ordering does not match (shuffled)
	for i := range testDeck {
		// Compare the two decks card by card.
		if testDeck[i] != unshuffledTestDeck[i] {
			// Found a card that didn't match, so break out of the loop.
			break
		}

		// Check if we have got to the end of the decks without finding a mismatch.
		if i == len(testDeck)-1 {
			t.Errorf("Expected at least one different card between the decks.")
		}
	}
}
