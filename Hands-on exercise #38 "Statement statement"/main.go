package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 { // Statement statement idiom is basically 2 statements, one after anothe - very similar to comma ok idiom
			fmt.Printf("Iteration %v, total count: %v and x is equal to %v\n", i, c, x)
			c++
		}
	}
}
