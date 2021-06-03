package main

import (
	"fmt"
	"log"
)

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	printEnglishGreeting(eb)

	sb := spanishBot{}

	printSpanishGreeting(sb)
}

// Note: If we don't use reference of receiver, then no need to mention the reference variable
func (englishBot) getGreeting() string {
	return fmt.Sprintf("Hi, %s !", "Raja")
}

func (spanishBot) getGreeting() string {
	return fmt.Sprintf("Hola, %s !", "Raja")
}

func printEnglishGreeting(eb englishBot) {
	log.Println(eb.getGreeting())
}

func printSpanishGreeting(sb spanishBot) {
	log.Println(sb.getGreeting())
}
