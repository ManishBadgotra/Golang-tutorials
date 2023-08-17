package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ageChecker(n int) {
	fmt.Print("Enter your Age: ")

	reader(numString)

	ageConv, err := strconv.ParseInt(strings.TrimSpace(numString), 10, 64)
	errorHandler(err)

	num = int(ageConv)

	if num < 18 {
		fmt.Println("You are Minor.")
	} else if num > 18 && num < 65 {
		fmt.Println("You are an Adult.")
	} else {
		fmt.Println("You are Senior Citizen.")
	}
}
