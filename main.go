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
}

func newCard() string {
	return "Five of Diamonds"
}
