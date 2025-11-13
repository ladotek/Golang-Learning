package main

import "fmt"

func main() {
	c := make(chan int, 2) // Makes a buffer channel, allows for certain values to sit in there - this allows 2 values to sit in there.

	c <- 42 // Send the value 42 into the channel
	c <- 43 // This is gonna run, cause we  have room for 2 values on the channel!

	fmt.Println(<-c) //Take the value (pull it) from the channel and print it
	fmt.Println(<-c) //Take the other value (pull it) from the channel and print it
}

/*
Channels are a synchronization primitive that allow Goroutines to communicate and synchronize execution.
You can send values into a channel and receive those values into another Goroutine. This allows you to coordinate tasks and send data safely between Goroutines.


https://go.dev/ref/spec#Channel_types

*/
