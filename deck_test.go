package main

import "testing"

// t.Errorf("Expected deck length of 16, but got %v", len(d))

func TestNewDeck(t *testing.T) {
	testCards := newDeck()

	if len(testCards) != 28 {
		t.Errorf("Expected deck length of 28, but got %v", len(testCards))
	}
}
