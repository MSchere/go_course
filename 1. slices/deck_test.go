package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 48 {
		t.Errorf("Invalid deck length: Expected length of 48 but got %v", len(deck))
	}
	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", deck[0])
	}
	if deck[len(deck)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', but got %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deck_test")
	deck := newDeck()
	deck.saveToFile("_deck_test")
	loadedDeck := newDeckFromFile("_deck_test")
	if len(loadedDeck) != 48 {
		t.Errorf("Invalid loaded deck length: Expected length of 48 but got %v", len(deck))
	}
	os.Remove("_deck_test")
}
