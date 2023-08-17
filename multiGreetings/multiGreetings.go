package main

import "fmt"

func multiGreetings(slice []string) {
	for _, name := range slice { // used ok comma syntax to ignore index output from Range (a)
		fmt.Printf("Namaste, %v\n", name)
	}
}
