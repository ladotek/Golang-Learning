package main

import "fmt"

func main() {
	x := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}) // x equals sum function, with parameters - ints inside {1, 2, 3, ...}
	fmt.Println(x)
}

func sum(ii []int) int { // Todd's updated version
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}

// Todd doesnt like name returns

// Named returns (named result parameters) are when you give names to the function’s return values in the function signature.

/*
This is the first version, the func sum higher is Todds version, thats more readable
func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}

*/
