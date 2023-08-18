package ifElse

import (
	"fmt"
	"strconv"
	"strings"
)

func PlusMinusChecker(n int) string {

	fmt.Print("Enter number to check for Positive, Negative or Zero: ")

	Reader(numString)

	num, err := strconv.ParseInt(strings.TrimSpace(numString), 10, 64)
	errorHandler(err)

	if num > 0 {
		s := fmt.Sprintf("%d is Positive number.\n", num)
		return s
	} else if num < 0 {
		s := fmt.Sprintf("%d is Negative number.\n", num)
		return s
	} else {
		s := fmt.Sprintf("%d is Zero.\n", num)
		return s
	}

}
