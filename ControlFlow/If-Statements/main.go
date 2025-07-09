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

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("Equal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("Equal to the meaning of life")
	} else {
		fmt.Println("Greater than the meaning of life")
	}

}
