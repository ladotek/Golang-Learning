package main

import "fmt"

func main() {
	x := foo()
	fmt.Println(x)

	y := bar()
	fmt.Println(y())

	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", bar) // Returns a function from a function
	fmt.Printf("%T\n", y)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

func foo() int {
	return 42
}

func bar() func() int {
	return func() int {
		return 43
	}
}
