package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // Embedding another struct
}

func main() {

	raja := person{
		firstName: "Raja",
		lastName:  "Lekkala",
		contact: contactInfo{
			email: "raja.lekkala@gmail.com",
			zip:   484848,
		},
	}

	fmt.Printf("%+v \n", raja) // {firstName:Raja lastName:Lekkala contact:{email:raja.lekkala@gmail.com zip:484848}}
}
