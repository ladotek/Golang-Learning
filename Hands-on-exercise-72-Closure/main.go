package main

import (
	"fmt"
	"math"
)

func main() {
	x := powinator(2.55)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powinator(a float64) func() float64 { // The "powinator" function takes in a, which is a float64, and returns a function that is going to return a float64
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c) // A to the power of c ("a" na potenco "c")
	}
}

/*
Closure is when we have “enclosed” the scope of a variable in some code block.
For this hands-on exercise, create a func which “encloses” the scope of a variable

*/
