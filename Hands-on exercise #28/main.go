package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ { // RUN PROGRAM 100 TIMES! - second part of Hands-on exercise #29
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("Iteration %v\t The value of x and y are %v and %v\t", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("x and y are both less than 4")
		case x > 6 && y > 6:
			fmt.Println("x and y are both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			println("None of the specified cases are met")
		}

	}
}
