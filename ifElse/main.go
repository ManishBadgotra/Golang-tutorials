package main

var (
	num, year              int
	email, pass, numString string
)

func main() {

	primeOrNot(num)

	plusMinusChecker(num)
	loginChecker(email, pass)

	oddEven(num)

	leapYearChecker(year)

	ageChecker(num)

}
