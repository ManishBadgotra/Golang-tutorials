package recursive

import (
	"fmt"
)

func IsPrime(n, i int) bool {
	if n <= 2 {
		return n == 2
	}
	if n%i == 0 {
		return false
	}
	if i*i > n {
		return true
	}
	return IsPrime(n, i+1)
}

func RecursiveForPrime() {
	fmt.Println("Prime numbers within the range of 1 to 100:")
	for i := 2; i <= 100; i++ {
		if IsPrime(i, 2) {
			fmt.Println(i)
		}
	}
}
