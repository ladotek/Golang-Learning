package main

import "fmt"

func main() {
	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()
}

/*

● Build and use an anonymous func
● anonymous func / func literal / lambda func
An anonymous func, aka a function literal or a lambda function, is a way to define a function
without giving it a name. Instead, you can directly declare and define a function inline wherever
it is needed. Anonymous funcs are commonly used when you want to pass a function as an
argument to another function or when you need to define a short-lived function for a specific
task.

*/
