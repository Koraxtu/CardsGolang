package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	cards.shuffle()
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

func newDeckFromFile(s string) deck {
	file, err := os.Open(s)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var hand string
	for scanner.Scan() {
		hand = scanner.Text()
	}
	arr := strings.Split(hand, ", ")
	return deck(arr)
}

func (d deck) shuffle() deck {
	for i, card := range d {
		randNum := rand.Intn(len(d))
		d[randNum], d[i] = card, d[randNum]
	}
	return d
}
