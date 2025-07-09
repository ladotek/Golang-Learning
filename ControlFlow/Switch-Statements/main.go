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

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	default:
		fmt.Println("This is the default case for x")
	}

	// No default fallthrough

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("Printed because of fallthrough: x is 41")
		fallthrough
	case 42:
		fmt.Println("Printed because of fallthrough: x is 42")
		fallthrough
	case 43:
		fmt.Println("Printed because of fallthrough: x is 43")
		fallthrough
	default:
		fmt.Println("Printed because of fallthrough: this is the default case for x")
	}
}
