package main

import "fmt"

// Create a new type of deck whcih is a silce of string

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
