package main

import "fmt"

func main() {
	x := outer()
	fmt.Println(x())
}

func outer() func() int { // Returns a function from a function
	return func() int {
		return 42
	}
}

/*

Create a func
	○ which returns a func
		■ which returns 42
● assign the returned func to a variable
● call the returned func
● print

*/
