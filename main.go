package main

import "fmt"

type contactInfo struct {
	email    string
	postCode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	nate := person{
		firstName: "Nate",
		lastName:  "Payne",
		contactInfo: contactInfo{
			email:    "email@real.com",
			postCode: "BS1 1AA",
		},
	}

	nate.updateFirstName("Nathaniel")
	nate.print()
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
