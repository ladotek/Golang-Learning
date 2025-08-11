package main

import "fmt"

func main() {
	//create an array
	x := [5]int{}

	//assign values to each index position
	for i := 0; i < 5; i++ {
		x[i] = i
	}

	// print out
	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}

}
