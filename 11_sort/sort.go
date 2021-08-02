package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func subGo(slice []int, c chan []int) {
	sort.Ints(slice)
	fmt.Print("\n subarray : ", slice)
	c <- slice
}

func groupBy(series []int) [][]int {
	var divided [][]int
	page := 4
	chunkSize := (len(series) + page - 1) / page

	for i := 0; i < len(series); i += chunkSize {
		end := i + chunkSize

		if end > len(series) {
			end = len(series)
		}

		divided = append(divided, series[i:end])
	}

	return divided
}

func convertToInt(slice []string) []int {
	var newSli []int
	for _, v := range slice {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			break
		}
		newSli = append(newSli, num)
	}
	return newSli
}

func main() {

	series := bufio.NewReader(os.Stdin)
	fmt.Printf("*********Seperate each integer by space and hit enter to process*****")
	fmt.Printf("\nEnter series (Atleast 4): ")
	line, err := series.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	line = strings.Trim(line, " \n")
	intSeries := convertToInt(strings.Split(line, " "))
	seriesLength := len(intSeries)
	if seriesLength < 4 {
		fmt.Print("Length of series doesn't have atleast 4 integers")
		os.Exit(3)
	}

	groupedSeries := groupBy(intSeries)
	channel1 := make(chan []int)
	channel2 := make(chan []int)
	channel3 := make(chan []int)
	channel4 := make(chan []int)
	go subGo(groupedSeries[0], channel1)
	set1 := <-channel1
	go subGo(groupedSeries[1], channel2)
	set2 := <-channel2
	go subGo(groupedSeries[2], channel3)
	set3 := <-channel3
	go subGo(groupedSeries[3], channel4)
	set4 := <-channel4
	firstMerge := append(set1, set2...)
	secondMerge := append(set3, set4...)
	finalMerge := append(firstMerge, secondMerge...)
	sort.Ints(finalMerge)
	fmt.Print("\nfinal : ", finalMerge)
}
