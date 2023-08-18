package forLoop

import "fmt"

func NestedLoop() {

	for i := 1; i <= 3; i++ {
		for j := 3; j >= 1; j-- {
			fmt.Println(i, j) // in this j will execute till its condition matched when i is 1 before incrementing i to 2.
		}
	}

}

func NestedLoopOnSlice() {

	fmt.Println("Using Nested for-loop on Multi-Dimensional Array.")

	var arr = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Println(arr[i][j])
		}
	}

}
