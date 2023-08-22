package goroutines

import (
	"fmt"
)

func Practice1() {

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
