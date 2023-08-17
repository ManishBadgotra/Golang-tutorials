package main

import "log"
// error handling function.
func errorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
