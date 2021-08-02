package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please enter a sting")
	str := strings.ToLower(readStringFromStdin())
	found := startsWith(str, "i") && endsWith(str, "n") && contains(str, "a")
	if found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}

func readStringFromStdin() string {
	var str string
	fmt.Scanln(&str) //here we ignore all kind of errors
	return str
}

func startsWith(str string, char string) bool {
	return strings.Index(str, char) == 0
}

func endsWith(str string, char string) bool {
	return strings.Index(str, char) == len(str)-1
}

func contains(str string, char string) bool {
	return strings.Contains(str, char)
}
