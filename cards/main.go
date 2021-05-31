package main

import "fmt"

var fileName string = "my_deck.txt"

func main() {
	cards := newDeck()
	cards.saveToFile(fileName)
	deck, _ := readFromFile(fileName)
	fmt.Println(deck)
}
