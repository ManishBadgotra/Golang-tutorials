package main

import "fmt"

func RecoverHandler() {
	recover := recover()

	if recover != nil {
		fmt.Println("RECOVERED:", recover)
	}
}
