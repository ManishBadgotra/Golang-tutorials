package forLoop

import "fmt"

var (
	slice []int
	maps  map[int]string
)

func RangeInFor() {
	/*
		Range in Go gives 2 output 1st is index and 2nd is value and can be used on Array, slice or map at a time.
		We can use "_" (underscore) to ignore one of any output in Range.
		Lets see 3 ways to use a Range on Slice and Map both.
	*/

	slice = []int{3, 5, 2, 6, 7, 0, 9, 8}

	fmt.Println("For Range on Slice starts from here.")
	fmt.Println(slice)
	for i, v := range slice {
		fmt.Println(i, v) // i prints the index of a value in slice on which it is present and v prints value that is inside the slice one by one.
	}

	for _, v := range slice {
		fmt.Println("Value in Slice", v) // This only prints value of a slice. For Index we only uses next example in which we want index only.
	}

	for i := range slice {
		fmt.Println("Index:", i) // By default if we only declares one variable then it gives the index of that slice or array.
	}

	fmt.Println("For Range on Map starts from here.")
	fmt.Println(maps)
	maps = map[int]string{1: "Hello", 2: "World", 3: "!!!"}

	for k, v := range maps {
		fmt.Printf("Key %v: Value %v\n", k, v) // k prints the key of a map and v prints value of a corresponding key in map.
	}

	for _, v := range maps {
		fmt.Println("Value of key:", v) // This only prints value of a map. For key we only uses next example in which we want key only.
	}

	for k := range maps {
		fmt.Println("Key:", k) // By default if we only declares one variable then it gives the key of that map.
	}
}
