package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sli := make([]int, 0, 3)

	for i := 1; i != 0; {
		var num string
		fmt.Printf("Enter an integer or X\n")
		fmt.Scan(&num)

		i = strings.Compare(num, "X")
		if i != 0 {
			iNum, _ := strconv.Atoi(num)

			sli = append(sli, iNum)
			sort.Ints(sli)

			fmt.Printf("The sorted slice: %v\n", sli)
		}
	}
}
