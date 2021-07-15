package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := newCard() // := only gets used for initial val
	fmt.Println(card)
	printState()
}

func newCard() string {
	return "Five of Diamonds"
}