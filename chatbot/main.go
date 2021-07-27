package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printEnglishGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printSpanishGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printEnglishGreeting(eb)
	printSpanishGreeting(sb)
}