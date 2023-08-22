package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(4)

	go fun1()
	go fun2()

	go practice2("Hello")
	go practice2("Here")

	wg.Wait()

}
func fun1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func1 ---> ", i)
		// time.Sleep(time.Duration(5 * time.Millisecond))
	}
	fmt.Println("Func1 completed.")
	wg.Done()
}
func fun2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func2 ---> ", i)
		// time.Sleep(time.Duration(6 * time.Millisecond))
	}
	fmt.Println("Func2 completed.")
	wg.Done()

}
