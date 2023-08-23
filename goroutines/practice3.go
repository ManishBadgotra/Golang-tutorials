package goroutines

import (
	"fmt"
	"log"
	"net/http"
)
func Practice3() {
	websiteList := []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://digixito.com",
		"https://maps.google.com",
	}

	for _, val := range websiteList {
		go GetStatusCode(val)
		// wg.Add(1)
	}

	wg.Wait()

}

func GetStatusCode(endpoint string) {
	defer wg.Done() // Gives response back when works of goroutine gets done.

	res, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
