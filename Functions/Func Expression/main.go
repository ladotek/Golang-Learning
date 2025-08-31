package main

import "fmt"

/* An expression is a combination of values, variables, operators, and function calls that are
evaluated to produce a single value. I
*/

func main() {
	foo()

	x := func() {
		fmt.Println("Anonymous func ran")
	}

	y := func(s string) {
		fmt.Println("This is an anonymous func showing my name: ", s)
	}

	x()
	y("Lado")
}

func foo() {
	fmt.Println("Foo ran")
}

// A named function:
//// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

//An anonymous function
// func(p parameter(s)) (r return() { code }
