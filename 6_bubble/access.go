package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(s []int, j int) {
	buffer := s[j]
	s[j] = s[j+1]
	s[j+1] = buffer
}

func BubbleSort(s []int) {
	l := len(s)
	for i := 0; i < l; i++ {
		for j := 0; j < (l - i - 1); j++ {
			if s[j] > s[j+1] {
				Swap(s, j)
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a sequence of up to 10 integers (space-separated):")
	text, _ := reader.ReadString('\n')
	ints := strings.Fields(text)
	s := make([]int, 0)
	for _, v := range ints {
		i, _ := strconv.Atoi(v)
		s = append(s, i)
	}
	BubbleSort(s)
	fmt.Println(s)
}
