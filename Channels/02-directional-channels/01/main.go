package main

import "fmt"

func main() {
	c := make(chan int, 2) //
	c <- 42                //
	c <- 43                //

	fmt.Println(<-c) //Take the value (pull it) from the channel and print it - CANT DO IT, BECAUSE ITS SEND ONLY; UNABLE TO RECEIVE!!
	fmt.Println(<-c)
	fmt.Println("--------------------------")
	fmt.Printf("%T\n", c) // Print type of channel

}
