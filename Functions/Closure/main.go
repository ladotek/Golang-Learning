package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4
	fmt.Println(f()) // 5
	fmt.Println(f()) // 6  --- Everytime its called, it reiterates

	g := incrementor() // Assigned to a different location "g", storing different values in different storage location
	fmt.Println(g())   // 1
	fmt.Println(g())   // 2
	fmt.Println(g())   // 3
	fmt.Println(g())   // 4
	fmt.Println(g())   // 5
	fmt.Println(g())   // 6
}

func incrementor() func() int {
	x := 0              // Declare and initialize a variable
	return func() int { // This function is "ENCLOSED" in the outer function incrementor -thats why its called closure
		x++
		return x
	}
}
