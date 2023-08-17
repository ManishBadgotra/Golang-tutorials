package main

import (
	"fmt"
	"strconv"
	"strings"
)

func oddEven(n int) string {

	fmt.Printf("Enter number to check for Even or Odd: ")

	reader(numString)

	numConv, err := strconv.ParseInt(strings.TrimSpace(numString), 10, 64)
	errorHandler(err)

	num = int(numConv)

	if num%2 == 0 {
		s := fmt.Sprintf("%d is Even.\n", num)
		return s
	} else {
		s := fmt.Sprintf("%d is Odd.\n", num)
		return s
	}
}
