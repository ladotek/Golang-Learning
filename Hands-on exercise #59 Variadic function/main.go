package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	yi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(foo(xi...))

	fmt.Println(bar(yi))

}

func foo(ii ...int) int {
	n := 0
	for _, v := range ii {
		n += v
	}
	return n
}

func bar(iii []int) int {
	n := 0
	for _, v := range iii {
		n += v
	}
	return n
}

/*
Hands-on exercise
● create a func with the identifier foo that
	○ takes in a variadic parameter of type int
	○ pass in a value of type []int into your func (unfurl the []int)
	○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
	○ takes in a parameter of type []int
	○ returns the sum of all values of type int passed in

*/
