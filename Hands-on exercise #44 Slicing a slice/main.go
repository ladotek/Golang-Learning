package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("%v\n", xi)
	fmt.Println("-------------------------------")

	fmt.Printf("%v\n", xi[:5])
	fmt.Println("------------------")

	fmt.Printf("%v\n", xi[5:])
	fmt.Println("------------------")

	fmt.Printf("%v\n", xi[2:7])
	fmt.Println("------------------")

	fmt.Printf("%v\n", xi[1:6])
	fmt.Println("------------------")

}

/*
Using the code from the previous example, use SLICING to create the following new slices
which are then printed:
● [42 43 44 45 46]
● [47 48 49 50 51]
● [44 45 46 47 48]
● [43 44 45 46 47]
*/
