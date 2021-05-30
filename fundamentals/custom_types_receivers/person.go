package main

import "fmt"

type Persons []string

func (p Persons) print() {
	for index, person := range p {
		fmt.Println(index, person)
	}
}
