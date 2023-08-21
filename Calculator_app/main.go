package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	x    float64
	y    float64
	char string
)

func main() {

	fmt.Println("Basic Calculator Program.")

	fmt.Print("Enter First value: ")
	input, err := Reader()

	errorHandler(err) // error handler

	input = strings.TrimSpace(input)
	floatInput, err := strconv.ParseFloat(input, 64)

	errorHandler(err)

	x = floatInput

	fmt.Print("Enter Operator (+,-,*,/): ")
	input, err = Reader()

	errorHandler(err)

	input = strings.TrimSpace(input)

	char = input

	fmt.Print("Enter Second value: ")
	input, err = Reader()

	errorHandler(err)

	input = strings.TrimSpace(input)
	floatInput, err = strconv.ParseFloat(input, 64)

	errorHandler(err)

	y = floatInput

	// fmt.Printf("%T , %T, %T\n", x, char, y)

	switch char {
	case "+":
		fmt.Printf("Answer: %.4f\n", x+y)
		time.Sleep(3 * time.Second)

	case "-":
		fmt.Printf("Answer: %.4f\n", x-y)
		time.Sleep(3 * time.Second)

	case "*":
		fmt.Printf("Answer: %.4f\n", x*y)
		time.Sleep(3 * time.Second)

	case "/":
		fmt.Printf("Answer: %.4f\n", x/y)
		time.Sleep(3 * time.Second)

	default:
		fmt.Println("You can only use Add(+), Subtract(-), Multiply(*), Divide(/) symbols.")
		fmt.Println("Try Again...")
		fmt.Println()
		main()
	}

	defer RecoverHandler()
}
