package main

import "fmt"

func main() {
	// 4!
	fmt.Println(4 * 3 * 2 * 1)

	fmt.Println(factorial(4))
	fmt.Println(factLoop(4))

}

func factorial(n int) int {
	fmt.Println("This is n", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//anything you can do with recursion, you can do with loops

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}

/*

factorial(4) ---------> 4 *
factorial(4-1) -------> 3 *
factorial(3-1) -------> 2 *
factorial(2-1) -------> 1 *
factorial(1-1) -------> 0 * // BASE CASE - this is where we want recursion to stop!

*/

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }(argument(s))

/*
The technique of solving a problem by breaking it
down into smaller subproblems of the same type. In other words, a recursive function
is a function that calls itself during its execution.

The function continues to call itself on smaller subproblems until it reaches a base
case, which is a simple case that can be solved directly without further recursion.
*/
