package main

import (
	"fmt"
	"log"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return fmt.Sprintf("Hi, %s !", "Raja")
}

func (spanishBot) getGreeting() string {
	return fmt.Sprintf("Hola, %s !", "Raja")
}

func printGreeting(b bot) {
	log.Println(b.getGreeting())
}
