package main

import "fmt"

// var txt string

func main() {
	fmt.Println("Iteration with i in for loop.")
	iteratorForLoop()

	println()
	println()
	fmt.Println("Print with for loop in while loop logic.")
	whileLoop()

	println()
	fmt.Println("Printing string with the help of loop")
	loopOverString("ankit")

	println()
	reverseForLoop()

	println()
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
