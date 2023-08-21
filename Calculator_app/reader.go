package main

import (
	"bufio"
	"os"
)

func Reader() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	return input, err
}
