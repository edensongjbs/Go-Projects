package main

import "fmt"

func main() {
	cards := newDeck()
	// cards = append(cards, "Six of Spades")
	cards.print()

	fmt.Println(cards.toString())
}

// func newCard() string {
// 	return "Five of Diamonds"
// }