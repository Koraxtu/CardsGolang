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

	savingDeck := hand.saveToFile("deckSave.txt")
	if savingDeck {
		// fmt.Println("Succes")
	} else {

	}
}

func newCard() string {
	return "Five of Diamonds"
}
