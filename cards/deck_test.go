package main

import "testing"

func TestCards(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected response include 16 cards %v", len(d))
	}
}
