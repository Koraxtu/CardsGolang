package main

import (
	"os"
	"testing"
)

// t.Errorf("Expected deck length of 16, but got %v", len(d))

func TestNewDeck(t *testing.T) {
	testCards := newDeck()

	if len(testCards) != 28 {
		t.Errorf("Expected deck length of 28, but got %v", len(testCards))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 28 {
		t.Errorf("Expected deck length of 28, but got %v", len(loadedDeck))

	}
	os.Remove("_decktesting")
}
