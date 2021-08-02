package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter the string (add double quotes to the input string): ")
	num, err := fmt.Scanf("%q", &str)
	if err != nil {
		fmt.Println("Err : ", err, num)
	}
	strInput := strings.ToLower(str)
	if strings.HasPrefix(strInput, "i") && strings.Contains(strInput, "a") && strings.HasSuffix(strInput, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
