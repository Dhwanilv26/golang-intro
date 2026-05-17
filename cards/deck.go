package main

import (
	"fmt"
	"os"
	"strings"
)

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

func (d deck) print() { // (d deck) is a receiver function for the type deck, to yeh cards.print() se hi chalega, as this is attached to type deck
	// receiver function pehle lega, then () output karega
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { // deal() () is a normal function, to yeh deal(cards,x) se hi chalega, this is not attached to any type, belongs to package

	// normal function pehle hi apna naam le lega
	return d[:handSize], d[handSize:]
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[:handSize]
}

func (d deck) toString() string { // converts ["red","yellow","blue"] into "red,yellow,blue"
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// everything is stored as a slice of bytes in the hard disk, as "ABC" ] =[65,66,67]
	return os.WriteFile(filename, []byte(d.toString()), 0666)
	// 0666 0 for octal rep, 6 for read and write permissions, 666 means owner, grp and others can r&w this file and not execute it

}
