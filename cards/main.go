package main

func main() {
	// cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()

	// remainingDeck.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// secondCards := newDeckFromFile("my_cards")

	// secondCards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
