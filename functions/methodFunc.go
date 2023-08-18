package functions

import "fmt"

var (
	y *int
	z int
)

// func SingleMethods(x int) {

// 	fmt.Println(y)

// }

func MultiMethods(x ...int) {

	for i := 0; i < len(x); i++ {
		y = &x[i]
		z = z + *y
	}

	fmt.Println(z)

}
