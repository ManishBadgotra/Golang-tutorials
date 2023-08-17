package functions

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Yoe       int
}

func (emp Employee) Fullname() {
	fmt.Printf("Hello, %v %v\nWow, you are %d years old.\n", emp.FirstName, emp.LastName, emp.Age)
}
