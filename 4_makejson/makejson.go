package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var personMap = make(map[string]string)
	name := bufio.NewReader(os.Stdin)
	address := bufio.NewReader(os.Stdin)
	fmt.Printf("\nEnter a name : ")
	nameline, err := name.ReadString('\n')
	nameline = strings.Trim(nameline, " \n")
	personMap["name"] = nameline
	fmt.Printf("\nEnter an address: ")
	addressline, err := address.ReadString('\n')
	addressline = strings.Trim(addressline, " \n")
	personMap["address"] = addressline

	jsonByte, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(jsonByte))
}
