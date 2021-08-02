package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

var ch_PhiloNum chan int

func (p Philo) eat(Philo_num int) {
	for i := 0; i < 3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat", Philo_num+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		fmt.Println("finishing eating", Philo_num+1)

	}

}

func host() {
	rand.Seed(time.Now().UnixNano())
	ch_PhiloNum = make(chan int, 2)
	//generate PhiloNumer
	for i := 0; i < 2; i++ {
		Philo_num := rand.Intn(5)
		ch_PhiloNum <- Philo_num
	}

}

func main() {

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		cs_num := rand.Intn(4)
		philos[i] = &Philo{CSticks[cs_num], CSticks[(cs_num+1)%5]}
	}

	for {
		fmt.Println("New routine")
		go host()
		for i := 0; i < 2; i++ {
			go func() {
				philo_num := <-ch_PhiloNum
				philos[philo_num].eat(philo_num)
			}()
		}
		fmt.Println("End routine")
	}
}
