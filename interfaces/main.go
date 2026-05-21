package main

import "fmt"

type bot interface { // interface type
	getGreeting() string
}

// getGreeting(string) string , not possible inside go, (go doesnt support method overloading), ofc not supported in structs too, and not supported in interfaces as well

type englishBot struct{} // concrete type
type spanishBot struct{}

// struct has to satisfy all methods of an interface to be considered that interface,
// interfaces in go can only have methods and not variables
// cant initialize or create any value of the interface types, can only do so from the concrete types (all go types like string, int and then the custom ones like englishbot) , interfaces in go are ofc abstract thats why
// interfaces are always implicit in go, we dont write englishbot 'satifies' bot interface
func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// can have very custom logic, though right now it is very easy
	return "Hello there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
