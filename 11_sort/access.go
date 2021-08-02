package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func sort_slice(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(slice)
	fmt.Print("Sorted slice: ")
	fmt.Println(slice)
}

func main() {
	slice := make([]int, 0, 3)
	for true {
		var next_num string
		fmt.Println("Enter the next integer (or X to stop):")
		fmt.Scan(&next_num)

		if next_num == "X" {
			fmt.Println("Beginning Sorting.....")
			break
		}

		next_int, _ := strconv.Atoi(next_num)
		slice = append(slice, next_int)
		fmt.Println("The array so far is:")
		fmt.Println(slice)
	}

	slice_len := len(slice)

	if slice_len < 4 {
		fmt.Println("Not enough integers to sort - please insert at least 4.")
		return
	}

	fmt.Println("Dividing Array in to 4 chunks")
	chunksize := len(slice) / 4
	chunk1 := slice[0:chunksize]
	chunk2 := slice[chunksize : chunksize*2]
	chunk3 := slice[chunksize*2 : chunksize*3]
	chunk4 := slice[chunksize*3:]

	var wg sync.WaitGroup
	wg.Add(4)

	go sort_slice(chunk1, &wg)
	go sort_slice(chunk2, &wg)
	go sort_slice(chunk3, &wg)
	go sort_slice(chunk4, &wg)

	wg.Wait()
	sort.Ints(slice)
	fmt.Print("The final sorted array is: ")
	fmt.Println(slice)
}
