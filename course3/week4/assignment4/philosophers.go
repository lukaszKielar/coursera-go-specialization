package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	id                            int
	leftChopstick, rightChopstick *Chopstick
	nMeals                        int
}

func (p *Philosopher) eat(c chan *Philosopher, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		if p.nMeals < 3 {
			c <- p

			fmt.Printf("starting to eat %d\n", p.id)
			p.leftChopstick.Lock()
			p.rightChopstick.Lock()

			p.nMeals++

			fmt.Printf("finishing eating %d\n", p.id)
			p.leftChopstick.Unlock()
			p.rightChopstick.Unlock()
		}
		wg.Done()
	}
}

func host(c chan *Philosopher) {
	for {
		if len(c) == 2 {
			<-c
			<-c

			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan *Philosopher, 2)

	var wg sync.WaitGroup
	wg.Add(15)

	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, chopsticks[i], chopsticks[(i+1)%5], 0}
	}

	go host(c)
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(c, &wg)
	}

	wg.Wait()
}
