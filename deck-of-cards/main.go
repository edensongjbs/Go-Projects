package main

import "fmt"

func main() {
	cards := newDeck()
	// cards = append(cards, "Six of Spades")
	cards.print()

<<<<<<< HEAD
	fmt.Println(cards.toString())
=======
	hand, remainingDeck := deal(cards, 5)
	fmt.Println(hand, remainingDeck)
>>>>>>> 802793cb82ced3deae69d38fddd3dcadf84fe768
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
