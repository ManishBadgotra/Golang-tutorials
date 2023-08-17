package recursive

import (
	"fmt"
)

func isPrime(n, i int) bool {
	if n <= 2 {
		return n == 2
	}
	if n%i == 0 {
		return false
	}
	if i*i > n {
		return true
	}
	return isPrime(n, i+1)
}

func recursiveForPrime() {
	fmt.Println("Prime numbers within the range of 1 to 100:")
	for i := 2; i <= 100; i++ {
		if isPrime(i, 2) {
			fmt.Println(i)
		}
	}
}
