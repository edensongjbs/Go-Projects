package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	//alex := person{"Alex", "Andersonm"}
	alex := person{firstName:"Alex", lastName:"Anderson"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}