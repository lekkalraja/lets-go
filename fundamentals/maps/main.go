package main

import "fmt"

func main() {
	countries := map[string]string{
		"IN":  "India",
		"SG":  "Singapore",
		"NZ":  "New Zealand",
		"ENG": "England",
	}

	fmt.Printf("[Countries] : %+v \n", countries)

	var states map[string]string
	fmt.Printf("[states] Declaring Map : %v \n", states)
	// states["AP"] = "Andhra Pradesh" // panic: assignment to entry in nil map

	states = make(map[string]string) // initializing map
	states["AP"] = "Andhra Pradesh"
	states["TN"] = "Tamil Nadu"
	states["TS"] = "Telangana State"
	fmt.Printf("states after adding values : %v \n", states)

	fmt.Printf("Full name of AP : %s \n", states["AP"])

	delete(states, "TS")

	fmt.Printf("states after deleting (TS) values : %v \n", states)

	print(countries)

}

// Iterating over map
func print(m map[string]string) {
	for key, value := range m {
		fmt.Printf("%v -> %v \n", key, value)
	}
}

/*
	RUN:
	===
	raja@raja-Latitude-3460:~/Documents/coding/golang/lets-go$ go run fundamentals/maps/main.go
	[Countries] : map[ENG:England IN:India NZ:New Zealand SG:Singapore]
	[states] Declaring Map : map[]
	states after adding values : map[AP:Andhra Pradesh TN:Tamil Nadu TS:Telangana State]
	Full name of AP : Andhra Pradesh
	states after deleting (TS) values : map[AP:Andhra Pradesh TN:Tamil Nadu]
	SG -> Singapore
	NZ -> New Zealand
	ENG -> England
	IN -> India
*/
