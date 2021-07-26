package main

import (
	"fmt"
	"strings"
)


type deck []string


func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {
		"Hearts",
		"Diamonds",
		"Spades",
		"Clubs",
	}

	cardValues := []string {
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

func (d deck) saveToFile() {

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}