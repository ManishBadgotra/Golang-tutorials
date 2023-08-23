package goroutines

import "fmt"

func Practice2(s string) {
	defer wg.Done() // Response of a function in the end of function end.

	greet := func(s string) {
		for i := 1; i < 6; i++ {
			fmt.Println(s, i)
		}
	}
	greet(s)
}
