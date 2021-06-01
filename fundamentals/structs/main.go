package main

import "fmt"

// Declaring struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	// 1 - Way of creating (initializing) struct
	janakiRam := person{"Janaki", "Ram"} // Based on order it will assign values

	// Another - Way of creating (initializing) struct
	seethaRam := person{
		firstName: "seetha",
		lastName:  "Ram",
	}

	fmt.Println(janakiRam)          // {Janaki Ram}
	fmt.Printf("%+v \n", seethaRam) // {firstName:seetha lastName:Ram}

	// Another - Way of creating (initializing) struct

	var radhaKrishna person
	fmt.Printf("%+v \n", radhaKrishna) // {firstName: lastName:} - Created struct with zero-values (default) of properties

	// Updating struct
	radhaKrishna.firstName = "Radha"
	radhaKrishna.lastName = "krishna"

	fmt.Printf("%+v \n", radhaKrishna) // {firstName:Radha lastName:krishna}
}
