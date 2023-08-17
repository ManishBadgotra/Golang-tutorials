package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner" // package for spinner
)

var (
	nameList []string = []string{}
	answer   bool     = true
)

func main() {

	fmt.Println("We welcomes you in Namaste Greetings Program.")

	time.Sleep(1 * time.Second) // program sleeps for 1 second before continuing.
	fmt.Println("Please wait...")

	/*spinner package usage started from here*/

	s := spinner.New(spinner.CharSets[9], 450*time.Millisecond) // Build our new spinner.
	s.Start()                                                   // Start the spinner.
	time.Sleep(5 * time.Second)                                 // Run for some time to simulate work.
	s.Stop()                                                    // Ends the spinner.

	/*spinner package usage ended here*/

	fmt.Println("Setting things up...")
	time.Sleep(1 * time.Second)

	for answer {
		fmt.Print("Enter the Name: ")
		reader := bufio.NewReader(os.Stdin)

		inputName, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		} else {
			nameList = append(nameList, inputName) // appends names given by user in slice to be used later.
		}

		fmt.Print("Do you want to enter one more name? [y/N] ") // takes input if user wants to continue adding more names.
		ansReader := bufio.NewReader(os.Stdin)
		ansInput, err := ansReader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
		} else {
			index := 0
			ansInput = strings.ToLower(ansInput)
			ansInput = string(GetChar(ansInput, index)) // changing Unicode into Strings.

			if ansInput == "y" {
				answer = true
			} else {
				answer = false
			}
		}

	}
	multiGreetings(nameList) // calling multiGreetings function in last.

}
