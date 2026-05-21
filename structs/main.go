package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// go assigns 0 / default values on empty structs
	// var alexp person
	// alex := person{firstName: "alex", lastName: "anderson"}
	// fmt.Printf("%+v", alexp)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "abc@gmail.com",
			zipCode: 360242,
		},
	}
	jim.updateName("jimmy")
	jim.print()

	x := 10
	updateInt(&x)
	fmt.Println(x)

}

func updateInt(x *int) {
	*x = 50
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
