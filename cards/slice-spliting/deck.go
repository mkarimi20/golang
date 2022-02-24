package main

import "fmt"

// Create a new type of deck whcih is a silce of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "hearts", "diamonds", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
