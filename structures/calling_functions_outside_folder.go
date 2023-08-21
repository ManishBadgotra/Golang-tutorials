package structures

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func (emp Employee) Fullname() {
	fmt.Printf("Hello, %v\n", emp.FirstName+" "+emp.LastName)
	fmt.Printf("Wow, you are %d years old.\n", emp.Age)
}
