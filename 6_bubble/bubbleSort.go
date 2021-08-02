package main

import (
	"fmt"
)

func promptUserInput() []int {
	var capacity int = 10

	var slice = make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		var num int
		fmt.Printf("\nEnter an integer to be sorted  : ")
		ptr, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Not a valid : ", err, ptr)
			break
		}
		slice = append(slice, num)
	}
	return slice
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		Swap(sli, i)
	}
}

func Swap(sli []int, i int) {
	for j := 0; j < len(sli)-i-1; j++ {
		if sli[j] > sli[j+1] {
			sli[j], sli[j+1] = sli[j+1], sli[j]
		}
	}
}

func main() {
	slice := promptUserInput()
	BubbleSort(slice)
	fmt.Println("*****Sorted*****")
	fmt.Println(slice)
}
