package main

import "fmt"

func loopOverString(s string) string {

	for i := 0; i < len(s); i++ {
	 // array gives UNICODE numbers by default. So we need to convert it in string with help of string() function.
		fmt.Print(string(s[i]), " ")
	}
	println()
	return s
}
