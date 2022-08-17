package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firsName string
	lastName string
	contactInfo
}

func main() {
	jim := person{
		firsName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email:   "jim.party@mail.com",
			zipCode: 71123,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firsName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
