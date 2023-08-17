package main

import "fmt"

func main() {
	empName := "Manish Sharma"
	ename := &empName

	fmt.Println(empName)

	empNameUpdate(ename)

	fmt.Println(*ename)
}

func empNameUpdate(ename *string) {
	*ename = "Ankit Sharma"
}
