package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(s string) {
		fmt.Println("This is an anonymous func showing my name: ", s)
	}("Lado")
}

func foo() {
	fmt.Println("Foo ran")
}

// A named function:
//// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

//An anonymous function
// func(p parameter(s)) (r return() { code }
