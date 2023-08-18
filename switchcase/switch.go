package switchCase

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Scase() {
	fmt.Print("Enter day of a week: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.ToLower(strings.TrimSpace(input))

	switch input {
	case "monday":
		fmt.Println("This is a Second day of the week.")
	case "tuesday":
		fmt.Println("This is a Third day of the week.")
	case "wednesday":
		fmt.Println("This is a Fourth day of the week.")
	case "thursday":
		fmt.Println("This is a Fifth day of the week.")
	case "friday":
		fmt.Println("This is a Sixth day of the week.")
	case "saturday":
		fmt.Println("This is a Seventh day of the week.")
	case "sunday":
		fmt.Println("This is a First day of the week.")
	default:
		fmt.Println("Oops! This is not a day of a week.")
	}

}
