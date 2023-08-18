package ifElse

import (
	"bufio"
	"os"
)

func Reader(t string) string {

	reader := bufio.NewReader(os.Stdin)

	txt, err := reader.ReadString('\n')

	errorHandler(err)

	return txt

}
