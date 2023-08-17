package main

import (
	"bufio"
	"os"
)

func reader(t string) string {

	reader := bufio.NewReader(os.Stdin)

	txt, err := reader.ReadString('\n')

	errorHandler(err)

	return txt

}
