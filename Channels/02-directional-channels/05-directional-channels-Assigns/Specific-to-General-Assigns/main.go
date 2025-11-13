package main

import "fmt"

func main() {
	c := make(chan int)    //
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("--------------------------")
	fmt.Printf("c\t%T\n", c)   // Print type of channel
	fmt.Printf("cr\t%T\n", cr) // Print type of channel
	fmt.Printf("cs\t%T\n", cs) // Print type of channel

	// specific to general doesn't assign!
	c = cr
	c = cs

}
