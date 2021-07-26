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

	cardSuits := []string{
		"Hearts",
		"Diamonds",
		"Spades",
		"Clubs",
	}

	cardValues := []string{
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

<<<<<<< HEAD
func (d deck) saveToFile() {

}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
=======
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
>>>>>>> 802793cb82ced3deae69d38fddd3dcadf84fe768
