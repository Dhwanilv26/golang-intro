package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// go assigns 0 / default values on empty structs 
	var alexp person
	alex := person{firstName: "alex", lastName: "anderson"}
	fmt.Printf("%+v",alexp)
	fmt.Printf("%+v",alex)
}
