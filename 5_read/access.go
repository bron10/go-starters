package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	fileName := askUserForFilename()
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	names := make([]Name, 0, 10)
	for scanner.Scan() {
		nameParts := strings.Split(scanner.Text(), " ")
		names = append(names, Name{fname: nameParts[0], lname: nameParts[1]})
	}

	for _, name := range names {
		fmt.Println(name.fname, name.lname)
	}

}

func askUserForFilename() string {
	fmt.Println("Enter file name")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimRight(str, "\n")
	return str
}
