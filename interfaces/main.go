package main

import "fmt"

type bot interface {
	getGreeting() string
}

// English bot
type englishBot struct{}

// Spanish bot
type spanishBot struct{}

func main() {
	// Create english and spanish bots.
	eb := englishBot{}
	sb := spanishBot{}

	// Print the greetings for the 2 bots.
	printGreeting(eb)
	printGreeting(sb)
}

// print greeting from the interface.  This will call the get greeting function
// for the specific type that implements getGreeting()
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// Greeting function for the english bot.
// NOTE: we are only defining the type on the
// receiver since we are not using a variable.
func (englishBot) getGreeting() string {
	return "Hello!"
}

// Greeting function for the spanish bot.
func (spanishBot) getGreeting() string {
	return "Hola!"
}
