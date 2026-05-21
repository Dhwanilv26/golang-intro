package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// in go, there is no support for method overloading, so cant have same function names in the same package
func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// can have very custom logic, though right now it is very easy
	return "Hello there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
