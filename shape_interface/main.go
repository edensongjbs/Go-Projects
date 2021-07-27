package main

import (
	"fmt"
	"math"
)

type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

type shape interface{
	GetArea() float64 
}

func (s square) GetArea() float64 {
	return math.Exp2(s.sideLength)
}

func (t triangle) GetArea() float64 {
	return (t.height*t.base)/2
}

func printArea(s shape) {
	fmt.Println(s.GetArea())
}

func main() {
	s := square{4}
	t := triangle{4, 3}
	printArea(s)
	printArea(t)
}