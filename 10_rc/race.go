package main

import (
	"fmt"
	"time"
)

/*
Race condition is when multiple threads are trying to access and manipulate the same variable.
In the code below, main, add1 and add2 are all accessing and changing the value of i.
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.
*/

func add1(pt *int) {
	(*pt)++
	fmt.Println(*pt)
}

func add2(pt *int) {
	(*pt) = *pt + 2
	fmt.Println(*pt)
}

func main() {
	i := 0
	time.Sleep(time.Second)
	go add1(&i)
	go add2(&i)
	i++
	fmt.Println("done", i)

}
