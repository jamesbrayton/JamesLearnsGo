package main

import "testing"

func TestEnglishBotGreeting(t *testing.T) {
	eb := englishBot{}

	if eb.getGreeting() != "Hello!" {
		t.Errorf("Expected %v, but got %v", "Hello!", eb.getGreeting())
	}
}

func TestSpanishBotGreeting(t *testing.T) {
	sb := spanishBot{}

	if sb.getGreeting() != "Hola!" {
		t.Errorf("Expected %v, but got %v", "Hola!", sb.getGreeting())
	}
}
