package main

func main() {
	//card := newCard()
	//fmt.Println(card)
	cards := deck{newCard(), "Ace of spades"}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"

}
