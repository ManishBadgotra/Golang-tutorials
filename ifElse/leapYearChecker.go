package ifElse

import (
	"fmt"
	"strconv"
	"strings"
)

func LeapYearChecker(y int) string {
	fmt.Print("Enter year to check for Leap year or not: ")

	Reader(numString)

	yearConv, err := strconv.ParseInt(strings.TrimSpace(numString), 10, 64)
	errorHandler(err)

	num = int(yearConv)

	if (year%400 == 0) || (year%100 != 0 && year%4 == 0) {
		s := fmt.Sprintf("%d is Leap Year.\n", year)
		return s
	} else {
		s := fmt.Sprintf("%d is not Leap Year.\n", year)
		return s
	}
}
