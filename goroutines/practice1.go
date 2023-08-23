package goroutines

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Practice1() {
	wg.Add(2) // Adding goroutine thread in the function

	go Fun1()
	go Fun2()

	wg.Wait()
}
func Fun1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func1 ---> ", i)
		// time.Sleep(time.Duration(5 * time.Millisecond))
	}
	fmt.Println("Func1 completed.")
	wg.Done()
}
func Fun2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func2 ---> ", i)
		// time.Sleep(time.Duration(6 * time.Millisecond))
	}
	fmt.Println("Func2 completed.")
	wg.Done()

}
