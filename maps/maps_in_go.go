package maps

import "fmt"

func MapsInGolang() {
	a := make(map[string]string)
	fmt.Println(a)

	a["JS"] = "JavaScript"
	a["PY"] = "Python"

	fmt.Println(a)
	fmt.Println(a["PY"])

	for key, value := range a{
		fmt.Printf("Key %v has value of %v\n", key, value)
	}
}
