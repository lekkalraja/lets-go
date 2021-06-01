package main

import "fmt"

/**
Notes:
=> `&variable` : Gives the memory address of the value this variable is pointing at (&raja -> 0xc0000c4000)
=> `*pointer` : Gives the value this memory address in pointing at (*pointerToRaja -> {firstName:Raja lastName:Lekkala})

== func (pointerToRaja *person) updateFirstName(newFirstName string)
=> *person (i.e. *(type)) => This is a type description - it means we're working with a pointer to a person
=> *pointerToRaja => This is an operator - it means we want to manipulate the value the pointer is referencing


=> We have 2 types of variables
	1) A variable which holds `address` => We can Turn `address` into `value` with *variable
	2) A variable which holds `value` => We can Turn `value` into `address` with &variable

*/

type person struct {
	firstName string
	lastName  string
}

func main() {
	raja := person{
		firstName: "Raja",
		lastName:  "Lekkala",
	}

	fmt.Printf("[Caller] Before Update : %+v  and Address : %p \n", raja, &raja) // Address : 0xc0000c4000
	rajaPointer := &raja
	rajaPointer.updateFirstName("Rajasekhar")
	fmt.Printf("[Caller] After Update : %+v and Address : %p \n", raja, &raja) // Address : 0xc0000c4000

	// Another way (~ shortcut) to call receiver function without need of address variable (raja person(type) -> can call pointerToRaja *person(type))
	raja.updateFirstName("Sekhar")
	fmt.Printf("[Caller] After Update : %+v and Address : %p \n", raja, &raja) // Address : 0xc0000c4000
}

func (pointerToRaja *person) updateFirstName(newFirstName string) {
	(*pointerToRaja).firstName = newFirstName
	fmt.Printf("[Updater] Updated : %+v and Address : %p \n", pointerToRaja, pointerToRaja) // Address : 0xc0000c4000
}

/*
	raja@raja-Latitude-3460:~/Documents/coding/golang/lets-go$ go run fundamentals/structs/pass_by_reference.go
	[Caller] Before Update : {firstName:Raja lastName:Lekkala}  and Address : 0xc0000c4000
	[Updater] Updated : &{firstName:Rajasekhar lastName:Lekkala} and Address : 0xc0000c4000
	[Caller] After Update : {firstName:Rajasekhar lastName:Lekkala} and Address : 0xc0000c4000
	[Updater] Updated : &{firstName:Sekhar lastName:Lekkala} and Address : 0xc0000c4000
	[Caller] After Update : {firstName:Sekhar lastName:Lekkala} and Address : 0xc0000c4000
*/
