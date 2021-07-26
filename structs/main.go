package main

import "fmt"

type contactInfo struct{
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	//alex := person{"Alex", "Andersonm"}
	alex := person{firstName:"Alex", lastName:"Anderson"}
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	jim := person{"Jim", "Party", contactInfo{"jimparty@gmail.com", 11029,},}
	jim.print()
	jimPointer := &jim
	jimPointer.updateName("jimmy")
	jim.print()
}