package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Fullname struct {
		fname string
		lname string
	}

	slice := make([]Fullname, 0)
	var filename string
	fmt.Printf("\nEnter file name : ")
	fmt.Scan(&filename)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		s := strings.Split(scanner.Text(), " ")
		var naming Fullname
		naming.fname = s[0]
		naming.lname = s[1]

		slice = append(slice, naming)

	}

	f.Close()

	for _, v := range slice {
		fmt.Println(v.fname, v.lname)
	}
}
