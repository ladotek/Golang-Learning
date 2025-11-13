package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)

	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// receive channel
func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v // Pull (fan) all values in from "even" channel into "fanin" channel
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v // Pull (fan) all values in from "odd" channel into "fanin" channel
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}

/*
FAN IN = Taking values from many channels, and putting those values onto one channel.

Rob Pike’s code
	https://play.golang.org/p/buy30qw5MM

*/
