package main

import "fmt"

func main() {

	cards := []string{"Five Of Diamonds", "Ace of Diamonds"}

	fmt.Println(cards)

	newSetOfCards := append(cards, "Six of Diamonds")
	fmt.Println(cards, newSetOfCards)

	for index, card := range cards {
		fmt.Println(index, card)
	}

	// If don't care of index varaible then ignore that by using `_`
	for _, card := range cards {
		fmt.Println(card)
	}
}
