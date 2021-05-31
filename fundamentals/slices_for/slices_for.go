package main

import (
	"fmt"
	"strings"
)

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

	// Byte Slice
	greet := "Hi, There"
	fmt.Println([]byte(greet)) // [72 105 44 32 84 104 101 114 101] -> Representation ofDecimal Format

	greets := []string{"Hi", "There"}
	//fmt.Println(([]byte)(greets)) // cannot convert []string to []byte
	greeting := strings.Join(greets, ",")
	fmt.Println([]byte(greeting)) // [72 105 44 84 104 101 114 101]
}
