package main

import "fmt"

func practice2(s string) {
	greet := func(s string) {
		for i := 1; i < 6; i++ {
			fmt.Println(s, i)
		}
	}
	greet(s)

	wg.Done()
}
