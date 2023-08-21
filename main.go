package main

import (
	"Go/interfaces"
	"Go/structures"
	"fmt"
)

func main() {

	// factorial := recursion.Factorial(5)
	// fmt.Println(factorial)

	// forLoop.NestedLoopOnSlice()

	/*
		Calling through Structs.
	*/

	emp1 := structures.Employee{
		FirstName: "Manish",
		LastName:  "Sharma",
		Age:       25,
	}

	emp1.Fullname()

	/*
		Calling through Interface.
	*/

	rect := interfaces.Rectangle{
		Length:  20,
		Breadth: 10,
	}

	fmt.Println("Rectangle:", rect.Parameter())
}
