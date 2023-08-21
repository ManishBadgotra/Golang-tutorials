package main

import "fmt"

var errors string

func errorHandler(err error) string {
	if err != nil {
		errors = fmt.Sprintln("Found error:", err.Error())
	}

	return errors
}
