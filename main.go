package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	vini := person{
		firstName: "Vinicius",
		lastName:  "Robledo",
		contactInfo: contactInfo{
			email:   "vini@gmail.com",
			zipCode: 123456,
		},
	}

	vini.print()

	vini.updateName("Amanda")

	vini.print()

}

func (p person) print() {
	//Printf tem modificadores como o +v que exibe nome dos atributos também (normalmente só printa os valores dos atributos)
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
