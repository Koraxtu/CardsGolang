package main

import "fmt"

func main() {
	cards := newDeck()
	// cards.print()

	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println("")
	remainingDeck.print()
	fmt.Println("")

	stringer := hand.toString()
	fmt.Println(stringer)

	savingDeck := remainingDeck.saveToFile("deckSave.txt")
	if savingDeck {
		// fmt.Println("Succes")
	} else {

	}
	// newestDeck := newDeckFromFile("deckSave.txt")
	// fmt.Println("")
	// newestDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
