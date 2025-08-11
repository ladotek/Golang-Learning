package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	b := a

	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("------------------")

	a[0] = 7

	fmt.Println("a", a)
	fmt.Println("b", b) //Because everything in a is also in b... Same printout. Both a and b are pointing to this underlying array
	fmt.Println("------------------")

}
