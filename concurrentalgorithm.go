package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	mut             *sync.Mutex
}

func (p Philo) eat(id int, request, finish chan int) {
	for i := 0; i < 3; i++ {
		// ask host
		request <- id

		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("starting to eat", id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		fmt.Println("finished eating", id)
		finish <- id

	}
	fmt.Println(">", id, "completed their 3 courses")
}

func host(wg *sync.WaitGroup, request, finish chan int) {
	defer wg.Done()
	eatingPeople := 0
	for i := 0; i < 3*5; i++ {
		select {
		case <-finish:
			eatingPeople--
		default:
			if eatingPeople >= 2 {
				fmt.Println(">", "Already have two people eating. Waiting...")
				r := <-finish
				eatingPeople--
				fmt.Println(">", r, "ended his meal. Continuing")
			}
		}
		<-request
		eatingPeople++
	}
}
func main() {

	philo := make([]*Philo, 5)
	CSticks := make([]*ChopS, 5)
	request := make(chan int, 2)
	finish := make(chan int, 2)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go host(wg, request, finish)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	mut := new(sync.Mutex)
	for i := 0; i < 5; i++ {
		philo[i] = &Philo{CSticks[i], CSticks[(i+1)%5], mut}
	}

	for i := 0; i < 5; i++ {
		go philo[i].eat(i+1, request, finish)
	}

	wg.Wait()

	fmt.Println("All philo completed their food")
}
