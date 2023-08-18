package main

import (
	"Go/forLoop"
	"Go/recursion"
	"fmt"
)

func main() {

	factorial := recursion.Factorial(5)
	fmt.Println(factorial)

	forLoop.NestedLoop()

}
