package main

import (
	"fmt"
	"strconv"
	"strings"
)

func primeOrNot(n int) string {
	fmt.Print("Enter number to check for Prime or not: ")
	reader(numString)

	num, err := strconv.ParseInt(strings.TrimSpace(numString), 10, 64)

	errorHandler(err)

	if (num%2 == 0) || (num%3 == 0) || (num%5 == 0) {

		s := fmt.Sprintf("%d is a prime number.\n\n", num)
		return s

	} else if (num%2 != 0) || (num%3 != 0) || (num%5 != 0) || (num%7 != 0) {

		s := fmt.Sprintf("%d is not a prime number.\n\n", num)
		return s

	} else {
		s := fmt.Sprintf("%d is Not a prime number.\n\n", num)
		return s

	}

}
