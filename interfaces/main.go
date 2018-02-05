package main

import "fmt"

// Bot interface.
// NOTE: Interfaces in Go are not generic types - in fact Go
//     famously does not have support for generic types.
// NOTE2: In Go, interfaces are implicit.
// NOTE3: Interfaces are a contract to help us manage types.  Also,
//     remember: garbage in -> garbage out.
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
