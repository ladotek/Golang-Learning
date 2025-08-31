package main

import "fmt"

// Callback = Passing a func as an argument

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(1, 2, add)
	fmt.Println(x)

	y := doMath(42, 16, subtract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int { // Takes in 2 values of int (a, b) and also takes in a function that takes in 2 values of int and returns an int
	return f(a, b)
}

func add(a int, b int) int { // Function "add" is gonna take a which is an int, and b, which is an int. Its gonna return an int.
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }
