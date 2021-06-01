package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	raja := person{
		firstName: "Raja",
		lastName:  "Lekkala",
	}

	fmt.Printf("[Caller] Before Update : %+v  and Address : %p \n", raja, &raja) // Address : 0xc000030040
	raja.updateFirstName("Rajasekhar")
	fmt.Printf("[Caller] After Update : %+v and Address : %p \n", raja, &raja) // Address : 0xc000030040
}

// Passing by value i.e. go internall fetch person(raja) value from original memory address
// and copy into some other available memory address and process the value
func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
	fmt.Printf("[Updater] Updated : %+v and Address : %p \n", p, &p) // Address : 0xc000030080
}

/*
	raja@raja-Latitude-3460:~/Documents/coding/golang/lets-go$ go run fundamentals/structs/pass_by_value.go
	[Caller] Before Update : {firstName:Raja lastName:Lekkala}  and Address : 0xc000030040
	[Updater] Updated : {firstName:Rajasekhar lastName:Lekkala} and Address : 0xc000030080
	[Caller] After Update : {firstName:Raja lastName:Lekkala} and Address : 0xc000030040
*/
