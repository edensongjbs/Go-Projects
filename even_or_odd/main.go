package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, n := range nums {
		var even_or_odd string
		if n%2 == 0 {
			even_or_odd = "even"
		} else {
			even_or_odd = "odd"
		}
		fmt.Println(n, " is ", even_or_odd)
	}
}