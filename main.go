package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	// append does not modify existing slices
	// returns a new slice which we are assigning back to the cards variable
	// preceding variable name could be a different variable
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
