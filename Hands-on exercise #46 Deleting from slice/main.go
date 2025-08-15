package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("%v\n", x)
	fmt.Println("-------------------------------")

	x = append(x[:3], x[5:]...)
	fmt.Printf("%v\n", x)

}

/*
To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise,
follow these steps:
● start with this slice
	○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
● use APPEND & SLICING to get these values here which you should ASSIGN to the
variable “x” and then print:
	○ [42, 43, 44, 48, 49, 50, 51]
*/
