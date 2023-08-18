package functions

var (
	y *int
	z int
)

func SingleMethods(x int) *int {
	/*This Function returns the memory address.*/
	return &x
}

func MultiMethods(x ...int) (string, int) {

	for i := 0; i < len(x); i++ {
		y = &x[i]
		z = z + *y
	}

	return "Total:", z

}
