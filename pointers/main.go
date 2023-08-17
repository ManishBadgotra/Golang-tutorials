package main

import "fmt"

var name string

func main() {
	empName := "Manish Sharma"
	ename := &empName

	fmt.Println(empName)

	fmt.Print("Enter name you want to enter: ")
	fmt.Scanln(&name)

	empNameUpdate(ename, name)

	fmt.Println(*ename)
}

func empNameUpdate(ptrname *string, text string) {
	*ptrname = text
}
