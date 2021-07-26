package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":"#ff0000",
		"green":"#00ff00",
		"blue":"#0000ff",
	}
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	delete(colors, "black")
	
	fmt.Println(colors)
}