package main

// import "fmt"

func main() {
	cards := newDeck()
	// cards.print()

// 	st, end := deal(cards, 4)
// 	st.print()
// 	fmt.Println("remaining cards")
// 	end.print()

// 	fmt.Println("NOW THE SAME WITH receiver FUNCTION")

// 	st, end = cards.deal(4)
// 	st.print()
// 	fmt.Println("remaining cards")
// 	end.print()

	// temp := "Hello world!"
	// fmt.Println([]byte(temp))

// 	fmt.Println(cards.toString())

// 	cards.saveToFile("temp_save")

// 	fmt.Println("printing deck from file:")
// 	fmt.Println(newDeckFromFile("temp_save"))

    cards.shuffle()
    cards.print()
	
	

}

func newCard() string {
	return "Five of Diamonds"
}
