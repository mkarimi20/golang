package main

import "testing"

func TestCards(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected response include 52 cards %v", len(d))
	}
	if d[0] != "Ace of spades" {
		t.Errorf("Expected Ace of spaced, but got %v", d[0])
	}

	if d[len(d)-1] != "King of clubs" {
		t.Errorf("Expected King of clubs, but got %v", d[len(d)-1])
	}
}
