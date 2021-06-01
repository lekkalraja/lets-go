package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // If we don't specify field name , then it will be same as type name i.e. contactInfo here
}

func main() {

	achillies := person{
		firstName: "Achillies",
		lastName:  "Warrior",
		contactInfo: contactInfo{
			email: "a@warrior.com",
			zip:   123454,
		},
	}

	achillies.sendMail()
	achillies.print()

}

func (p person) sendMail() {
	fmt.Printf("%s requested to send mail for id : %s \n", p.firstName, p.contactInfo.email)
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

/*
TO RUN :
========
raja@raja-Latitude-3460:~/Documents/coding/golang/lets-go$ go run fundamentals/structs/structs_receiver.go
Achillies requested to send mail for id : a@warrior.com
{firstName:Achillies lastName:Warrior contactInfo:{email:a@warrior.com zip:123454}}
*/
