package main

import "fmt"

func main() {
	//cards := newDeck()
	cards := newDeckFromFile("test-deck.txt")
	cards.shuffle()
	// cards = append(cards, "Six of Spades")
	// cards.print()

	// fmt.Println(cards.toString())
	// cards.saveToFile("test-deck.txt")
	hand, remainingDeck := deal(cards, 5)
	fmt.Println(hand, remainingDeck)
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
