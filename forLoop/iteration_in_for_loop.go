package forLoop

import "fmt"

func IteratorForLoop() {
	for i := 1; i <= 10; i++ {
		if i == 10 {
			fmt.Print(i)
		} else {
			fmt.Print(i, ",")
		}

	}
}
