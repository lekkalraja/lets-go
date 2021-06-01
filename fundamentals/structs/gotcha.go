package main

import "fmt"

func main() {
	greet := []string{"Hi", "How", "are", "you?"}

	fmt.Println(greet) // [Hi How are you?]

	updateSlice(greet)

	fmt.Println(greet) // [Bye How are you?]  -> updated slice without using pointers
}

func updateSlice(greet []string) {
	greet[0] = "Bye"
}

/*
	Note:
	=====
	* Arrays -> Primitive data structure, Can't be resized. Rarely used directly
	* Slices -> Can grow and shrink, Used 99% of the time for lists of elements

	  greet := []string{"Hi", "How", "are", "you?"} -> Internally it will be like below

	  greet

	  |slice|
*/
