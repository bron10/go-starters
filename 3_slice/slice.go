package main

import (
	"fmt"
	"sort"
)

func main() {
	var sli = make([]int, 0, 3)

	for {
		var num int
		fmt.Printf("Enter an integer to be added to the slice: ")
		ptr, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Err : ", err, ptr)
			break
		}
		sli = append(sli, num)
		sort.Ints(sli)
		fmt.Printf("\nPrint slice ", sli)
	}
}
