package main

import "fmt"

func main() {
	foo()
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }(argument(s))

func foo() {
	fmt.Println("foo ran")
}
