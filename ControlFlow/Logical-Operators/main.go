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

	//CONDITIONAL
	// if statements
	// switch statements

	if x < 42 && y < 42 {
		//// requires both to be true to run
		fmt.Println("Both are less than the meaning of life.")
	}

}
