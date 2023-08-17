package structFolder

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Calling() {
	fmt.Println("Hello, World!!!")

	var p1 Person = Person{"Bheem Singh", 21}

	fmt.Println(p1)
}
