package main

import (
	"fmt"
)

func promptUserInput(text string) float64 {
	var num float64
	fmt.Printf("\nEnter  %s : ", text)
	ptr, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Not a valid : ", err, ptr)
	}
	return num
}

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 1/2*a*t*t + vo*t + so
	}
	return fn
}

func main() {

	acceleration := promptUserInput("Acceleration")
	initVelocity := promptUserInput("Initial Velocity")
	initDisplacement := promptUserInput("Initial displacement")
	displaceBytime := GenDisplaceFn(acceleration, initVelocity, initDisplacement)
	time := promptUserInput("Time in seconds")
	fmt.Printf("Printing displacement with entered time %f secs: ", time)
	fmt.Print(displaceBytime(time))
	fmt.Printf("\n ***Example Printing displacement for 3 and 5 seconds respectively***")

	fmt.Printf("\n %f , %f", displaceBytime(3), displaceBytime(5))
}
