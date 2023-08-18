package functions

import "fmt"

var counter int

func Anonymous() {
	counter = 0

	func() {
		fmt.Println("This Anonymous function will only run when this Anonymous() is called First Time.")
	}()

	varFunction := func() {
		counter += 1

		if counter == 1 {
			fmt.Printf("This Anonymous function is called %v time through Variable and can be called as many times as we want.\n", counter)

		} else {

			fmt.Printf("This Anonymous function is called %v times through Variable and can be called as many times as we want.\n", counter)

		}

	}

	varFunction()
	varFunction()
	varFunction()
	varFunction()
	varFunction()
	varFunction()
	varFunction()
	varFunction()

}
