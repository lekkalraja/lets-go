package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("Somethign went wrong while opening the file : %v \n", err)
	}

	io.Copy(os.Stdout, file)

	/*
		OUTPUT:
		raja@raja-Latitude-3460:~/Documents/coding/golang/lets-go$ go run fundamentals/interfaces/assignment_2.go README.md
		# LET's GO
	*/

}

/*
	Hard Mode Interfaces:
	* Create a program that reads the contents of a text file tehn prints it's contents to the terminal.
	  The file to open should be provided as an argument to the program when it is executed at the terminal.
	  	ex : go run main.go myfile.txt should open up the myfile.txt

	What interfaces does the `File` type implement ?

	If the `File` type implements the `Reader` interace, you might be able to reuse that io.Copy function!
*/
