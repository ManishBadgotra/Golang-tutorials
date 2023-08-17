package main

import (
	"fmt"
	"os"
)

func readString(filename string) {
	data, err := os.ReadFile(filename)

	errorHandler(err)

	fmt.Println(string(data))
}
