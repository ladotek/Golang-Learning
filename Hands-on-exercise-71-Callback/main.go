package main

import "fmt"

func main() {
	fmt.Println(printSquare(square, 3)) // Function printSquare will get passed in "square" and also 3...

}

func printSquare(f func(int) int, a int) string { // Function "printSquare" takes in "f which is a function that is going to take in an int and return an int", and "a which is an int", and will return a string
	x := f(a) // Calls the function f with argument a. With f == square and a == 3, x becomes 9
	return fmt.Sprintf("the number %v squared is %v", a, x)
}

func square(n int) int { // Function "square" takes in "n which is an int", and returns and int.
	return n * n
}

/*
Call flow (at runtime)

1. main → printSquare(square, 3)
2. Inside printSquare: call f(a) → square(3) → returns 9
3. Format the message → return "the number 3 squared is 9"
4. fmt.Println(...) prints

*/

/*
A “callback” is when we pass a func into a func as an argument. For this exercise,
● pass a func into a func as an argument
	○ func square(n int) int
	○ printSquare(f func(int)int, int) string

*/
