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

	cardValues := []string{"Ace", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}
	index2Value := cardValues[2] //2

	first5Elements := cardValues[0:5]
	first5Elements2 := cardValues[:5]
	after5thElement := cardValues[5:len(cardValues)]
	after5thElement2 := cardValues[5:]
	fmt.Println(index2Value, first5Elements, first5Elements2, after5thElement, after5thElement2)
}
