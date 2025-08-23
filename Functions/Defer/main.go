package main

import "fmt"

func main() {
	defer foo() // defer = dont run this function until the function surrounding it exits! defer f.Close (closing a file example...)
	bar()

}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
