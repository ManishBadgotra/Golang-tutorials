package pointer

import "fmt"

var name string

func PointerWithStruct() {
	empName := "Manish Sharma" // initilize and declare variable.
	ename := &empName          // initilize and declares memory address of empName.

	fmt.Println(empName)

	fmt.Print("Enter name you want to enter: ")
	fmt.Scanln(&name) // takes string and stores it in "name" variable

	empNameUpdate(ename, name)

	fmt.Println(*ename) //prints ename variable's value by adding * before variable name.
}

func empNameUpdate(ptrname *string, text string) {
	*ptrname = text // takes text from user and updates it to pointer variable.
}
