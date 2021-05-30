package main

import "fmt"

var author string = "Raja"
var year int

// place := "Singapore" // Invalid Declaration

func main() {
	var card string = "Ace of Cards"
	fmt.Println(card)

	totalNumberOfCards := 52
	fmt.Println(totalNumberOfCards)

	// totalNumberOfCards := 53 // ERROR : no new variables on left side of :=

	totalNumberOfCards = 53
	fmt.Println(totalNumberOfCards)

	// Assigning to a global variable
	year = 2021

	fmt.Println(author, year)
}
