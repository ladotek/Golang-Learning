package main

import "fmt"

func main() {
	c := make(<-chan int, 2) // This channel is a receive-only channel!

	c <- 42 // CANT SEND TO IT!!  cannot send to receive-only channel c
	c <- 43 // CANT SEND TO IT!!  cannot send to receive-only channel c

	fmt.Println(<-c) //Take the value (pull it) from the channel and print it ; CANT DO IT, BECAUSE ITS RECEIVE ONLY!
	fmt.Println(<-c)
	fmt.Println("--------------------------")
	fmt.Printf("%T\n", c) // Print type of channel

}
