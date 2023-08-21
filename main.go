package main

import (
	"Go/interfaces"
	"fmt"
)

func main() {

	// factorial := recursion.Factorial(5)
	// fmt.Println(factorial)

	// forLoop.NestedLoopOnSlice()

	/*
		Calling through Structs.
	*/

	// emp1 := structures.Employee{
	// 	FirstName: "Manish",
	// 	LastName:  "Sharma",
	// 	Age:       25,
	// }

	// emp1.Fullname()

	/*
		Calling area_and_perimeter Interface.
	*/

	// rect := interfaces.Rectangle{
	// 	Length:  20,
	// 	Breadth: 10,
	// }

	// fmt.Println("Rectangle:", rect.Parameter())

	/*
		Calling animal via Interfaces.
	*/

	var ketu = interfaces.Dog{
		Animal: interfaces.Animal{
			Breed:   "Canine",
			Sound:   "Woof",
			Nature:  "Domestic",
			CanFly:  false,
			CanSwim: true,
			CanWalk: true,
		},
		Name: "Ketu",
		Legs: 4,
	}

	fmt.Println(ketu)
}
