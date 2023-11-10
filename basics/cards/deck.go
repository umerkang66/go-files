package main

import "fmt"

// create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	// don't put receiver on this function
	cards := deck{}

	cardsSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardsSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}

	return cards
}

// Receiver function
func (d deck) print() {
	// this (d deck) is receiver
	for i, card := range d {
		fmt.Println(i, card)
	}
}
