package ifElse

import (
	"fmt"
	"strings"
)

func LoginChecker(email string, pass string) {

	fmt.Print("Enter your E-mail ID: ")
	Reader(email) // getting an input from user.
	email = strings.TrimSpace(email)

	fmt.Print("Enter your Password: ")
	Reader(pass)
	pass = strings.TrimSpace(pass)

	if email == "manish@go.dev" && pass == "manish123" {
		fmt.Println("Login Successfull.")
		println() // for single line space

	} else {
		fmt.Println("Login Unsuccessfull.")
		println()
	}
}
