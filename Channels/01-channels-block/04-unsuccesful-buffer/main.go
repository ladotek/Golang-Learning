package main

import "fmt"

func main() {
	c := make(chan int, 1) // Makes a buffer channel, allows for certain values to sit in there - this allows 1 to sit in there.

	c <- 42 // Send the value 42 into the channel
	c <- 43 // This isnt gonna run, cause we only have room for 1 value on the channel!

	fmt.Println(<-c) //Receive the value from the channel and print it

}

/*
Channels are a synchronization primitive that allow Goroutines to communicate and synchronize execution.
You can send values into a channel and receive those values into another Goroutine. This allows you to coordinate tasks and send data safely between Goroutines.

*/
