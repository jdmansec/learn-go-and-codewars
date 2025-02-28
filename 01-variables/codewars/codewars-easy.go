package main

import (
	"fmt"
	"strings"
)

// MakeUpperCase (8 kyu): Write a function which converts the input string to uppercase
func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}

func main() {
	// MakeUpperCase (8 kyu)
	fmt.Println(MakeUpperCase("james"))
}
