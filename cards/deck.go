package main

import "fmt"

// create a new type of 'deck' which is a slice of strings

type deck []string // deck extends string (sort of)

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards

}

func (d deck) print() { // (d deck) is a receiver function for the type deck
	for i, card := range d {
		fmt.Println(i, card)
	}
}
