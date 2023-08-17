package structFolder

import "fmt"

type Person struct {
	name string
	age  int
}

func calling() {
	fmt.Println("Hello, World!!!")

	var p1 Person = Person{"Ashwani", 21}

	fmt.Println(p1)
}
