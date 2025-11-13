package main

import "fmt"

func main() {
	c := make(chan int) // Makes a channel of type int

	go func() { // This goroutine is being sent to run on its own, the main goroutine flow continues!
		c <- 42 // Send the value 42 into the channel

	}()

	fmt.Println(<-c) //Receive the value from the channel and print it

}

/*
Channels are a synchronization primitive that allow Goroutines to communicate and synchronize execution.
You can send values into a channel and receive those values into another Goroutine. This allows you to coordinate tasks and send data safely between Goroutines.

*/
