package main

import "fmt"

func main() {
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	card := "Five of Diamonds"
	return card
}
