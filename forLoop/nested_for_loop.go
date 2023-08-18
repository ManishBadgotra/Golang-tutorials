package forLoop

import "fmt"

func NestedLoop() {

	for i := 1; i <= 3; i++ {
		for j := 3; j >= 1; j-- {
			fmt.Println(i, j)
		}
	}
}
