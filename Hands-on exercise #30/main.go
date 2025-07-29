package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("Iteration %v\t", i)

		switch {
		case x == 0:
			fmt.Printf("The value of x is %v\n", x)
		case x == 1:
			fmt.Printf("The value of x is %v\n", x)
		case x == 2:
			fmt.Printf("The value of x is %v\n", x)
		case x == 3:
			fmt.Printf("The value of x is %v\n", x)
		case x == 4:
			fmt.Printf("The value of x is %v\n", x)
		}
		//Another option, cleaner!
		/*switch x {
		case 0:
			fmt.Printf("Iteration %v \t The value of x is %v\n", i, x)
		case 1:
			fmt.Printf("Iteration %v \t The value of x is %v\n", i, x)
		case 2:
			fmt.Printf("Iteration %v \t The value of x is %v\n", i, x)
		case 3:
			fmt.Printf("Iteration %v \t The value of x is %v\n", i, x)
		case 4:
			fmt.Printf("Iteration %v \t The value of x is %v\n", i, x)
		}
		*/

	}
}
