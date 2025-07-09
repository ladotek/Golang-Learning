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

	// LOOPS / INTERACTIVE
	// for statements

	for i := 0; i < 5; i++ {
		fmt.Printf("Counting to five: %v \t first for loop\n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop\n", y)
		y++
	}

	// break
	// takes you out of the loop

	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue
	// takes to next iteration

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Counting even numbers: ", i)
	}
}
