package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is the first statement to run")
	fmt.Println("This is the second statement to run")
	x := 40 //this is the third statement to run
	y := 5  //this is the fourth statement to run
	fmt.Printf("c=%v \n y=%v\n", x, y)
	//Switch statements - (SWITCH - CASE)

	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x is equal to 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("This is the default case for x")
	}

	// A "select" statement chooses which of a set of possible send or receive operations will proceed.
	// It looks similar to a "switch" statement but with the cases all referring to communication operations.
	//https://go.dev/ref/spec#Select_statements

	/*select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)

	}
	*/
}
