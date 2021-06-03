package main

import (
	"log"
	"reflect"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	s := square{
		sideLength: 4.0,
	}

	printArea(s)

	t := triangle{
		height: 3.0,
		base:   4.0,
	}

	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	switch reflect.TypeOf(s) {
	case reflect.TypeOf(triangle{}):
		log.Printf("The triangle has an area of %.2f\n", s.getArea())
		break
	case reflect.TypeOf(square{}):
		log.Printf("The square has an area of %.2f\n", s.getArea())
	}
}

/*
	Problem:
	========
	Write a program that creates two custom struct types called "triangle" & "square"

	The "square" type should be a struct with a filed called "sideLength" of type "float64"
	The "triangle" type should be a struct with a field called "height" of type "float64"
	and a field "base" of type "float64"

	Both types should have function called "getArea"
	that returns the calculated area of the square or triangle

	Area of a triangle = 0.5 * base * height
	Area of a square = sideLength * sideLength

	Add a "shape" interface that defines a function called "printArea".
	This function should calculate the area of the given shape and
	print it out to the terminal design the interface
	so that the "printArea" function can be called with either a triangle or a square.

	RUN :
	=====
	raja@raja-Latitude-3460:~/Documents/coding/golang/lets-go$ go run fundamentals/interfaces/assignment_1.go
	2021/06/03 11:46:42 The square has an area of 16.00
	2021/06/03 11:46:42 The triangle has an area of 6.00
*/
