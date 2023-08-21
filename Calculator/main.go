package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	x    float64
	y    float64
	char rune
)

func main() {
	fmt.Println("Basic Calculator Program.")

	fmt.Print("Enter First value: ")
	input, err := Reader()

	if err != nil {
		fmt.Println(err.Error())
	}

	input = strings.TrimSpace(input)
	floatInput, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Found Error in:", err)
	}

	x = floatInput

	fmt.Print("Enter Operator: ")
	input, err = Reader()

	if err != nil {
		fmt.Println(err.Error())
	}

	input = strings.TrimSpace(input)

	for _, runeData := range input {
		char = runeData
	}

	fmt.Print("Enter Second value: ")
	input, err = Reader()

	if err != nil {
		fmt.Println("Found Error in:", err)
	}

	input = strings.TrimSpace(input)
	floatInput, err = strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println(err.Error())
	}

	y = floatInput

	switch char {
	case '+':
		fmt.Printf("Your Total is %.4f\n", x+y)
	case '-':
		fmt.Printf("Your Total is %.4f\n", x-y)
	case '*':
		fmt.Printf("Your Total is %.4f\n", x*y)
	case '/':
		fmt.Printf("Your Total is %.4f\n", x/y)
	default:
		fmt.Println("You can only use Add(+), Subtract(-), Multiply(*), Divide(/) symbols.")
	}
}
