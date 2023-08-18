package recursive

func Factorial(n int) int {
	if n == 0 {
		return int(1)
	} else {
		return n * Factorial(n-1)
	}
}
