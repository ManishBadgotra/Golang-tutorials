package main

// returns Unicode code point of 0th index in string.
func GetChar(str string, index int) rune {

	return []rune(str)[index] // makes rune slice of str and picks only 0th index value from it.
}
