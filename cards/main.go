package main

var fileName string = "my_deck.txt"

func main() {
	cards := newDeck()
	/*cards.saveToFile(fileName)
	deck, _ := readFromFile(fileName)*/
	cards.shuffle()
	cards.print()
}
