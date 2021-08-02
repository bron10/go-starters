package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	num             int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(c chan int, wg *sync.WaitGroup) {
	//Each philosopher should eat only 3 times
	c <- p.num
	for i := 1; i <= 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("\nstarting to eat ", p.num)
		fmt.Println("\nfinishing eating ", p.num)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		if i == 3 {
			wg.Done()
		}
	}
}

//a philosopher must get permission from a host which executes in its own goroutine
func host(c chan int, wg *sync.WaitGroup) {
	for {
		//host allows no more than 2 philosophers to eat concurrently
		if len(c) == 2 {
			phil1 := <-c
			phil2 := <-c
			//time delay
			fmt.Print("\n Eating phil no :", phil1)
			fmt.Print("\n Eating phil no : ", phil2)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 2)
	CSticks := make([]*ChopS, 5)
	wg.Add(15)

	//There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)

	//The philosophers pick up the chopsticks in any order, not lowest-numbered first
	for i := 4; i >= 0; i-- {
		n := 0
		if i == 0 {
			n = 4
		} else {
			n = i - 1
		}
		philos[i] = &Philo{
			//Each philosopher is numbered, 1 through 5.
			i + 1,
			CSticks[i],
			CSticks[n],
		}
	}

	go host(c, &wg)
	for i := 4; i >= 0; i-- {
		go philos[i].eat(c, &wg)
	}
	wg.Wait()
}
