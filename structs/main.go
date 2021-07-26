package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	//alex := person{"Alex", "Andersonm"}
	alex := person{firstName:"Alex", lastName:"Anderson"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	jim := person{"Jim", "Party", contactInfo{"jimparty@gmail.com", 11029,},}
	fmt.Printf("%+v", jim)
}