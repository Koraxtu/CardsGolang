package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of deck
// Which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Jack", "Queen"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	arr := []string(d)
	str := strings.Join(arr, ", ")
	return str
}

func (d deck) saveToFile(s string) bool {
	saveValue := d.toString()
	file, err := os.Create(s)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return false
	}
	defer file.Close()

	_, err = file.WriteString(saveValue)
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return false
	}
	return true
}
