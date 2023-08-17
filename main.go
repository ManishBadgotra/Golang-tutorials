package main

import (
	"Go/functions"
	"strconv"
)

var (
	// firstName string
	// lastName  string
	// age       int
	stringInt int64
	data      []string
)

func main() {
	// fmt.Print("Enter First name: ")
	// fmt.Scanln(&firstName)

	// fmt.Print("Enter Last name: ")
	// fmt.Scanln(&lastName)

	// fmt.Print("Enter age: ")
	// fmt.Scanln(&age)

	data = []string{"Manish", "Sharma", "35"}
	stringInt, _ = strconv.ParseInt(data[2], 10, 64)

	e1 := functions.Employee{FirstName: data[0], LastName: data[1], Age: int(stringInt)}

	e1.Fullname()
}
