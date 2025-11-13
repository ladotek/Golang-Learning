package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go foo(c)

	// receive
	bar(c) // Initially this was: go bar(c) --- Removing the "go" before it is essential, as it needs to wait for go foo(c) to finish sending c, before bar(c) can pull it off!

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
