package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	fmt.Print("This Program Read/Write from File.\n")

	spin := spinner.New(spinner.CharSets[9], 300*time.Millisecond)
	spin.Start()
	time.Sleep(3 * time.Second)
	spin.Stop()

	content := "This line is entered with the help of Go program.\nThis is a second line pushed into the file.\nThis is a last line entered in the file."

	file, err := os.Create("./plainTextFile.txt")

	errorHandler(err) // error handling through error handling function.

	_, err = io.WriteString(file, content)
	errorHandler(err)

	readString("./plainTextFile.txt")

	defer file.Close()
}
