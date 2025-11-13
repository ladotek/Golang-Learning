package main

import "fmt"

func main() {
	c := make(chan<- int, 2) // This channel is a send-only channel!

	c <- 42 //
	c <- 43 //

	fmt.Println(<-c) //Take the value (pull it) from the channel and print it -- DOES NOT WORK, BECAUSE YOU CANT PULL FROM A SEND-ONLY CHANNEL!!
	fmt.Println(<-c)
	fmt.Println("--------------------------")
	fmt.Printf("%T\n", c) // Print type of channel

}
