package main

import "testing"

func TestCards(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected response include 52 cards %v", len(d))
	}
}
