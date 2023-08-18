package functions

import "fmt"

var y *int 

func SingleMethods(x int) {

	y = &x

	fmt.Println(*y)

}

func MultiMethods(x ...int) {

	for i := 0; i < len(x); i++ {
		y =  &x[i]
		*y = *y + *y
	}

	fmt.Println(*y)

}
