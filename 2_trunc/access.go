package main

import (
	"fmt"
)

func main() {

	var x float32
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&x)
	fmt.Printf("Truncated %f = %d", x, int(x))
}
