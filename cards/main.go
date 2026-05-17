package main

import "fmt"

func main() {
	cards := newDeck()
	// cards.print()

	st, end := deal(cards, 4)
	st.print()
	fmt.Println("remaining cards")
	end.print()

	fmt.Println("NOW THE SAME WITH receiver FUNCTION")

	st, end = cards.deal(4)
	st.print()
	fmt.Println("remaining cards")
	end.print()

	// temp := "Hello world!"
	// fmt.Println([]byte(temp))

	fmt.Println(cards.toString())

}

func newCard() string {
	return "Five of Diamonds"
}
