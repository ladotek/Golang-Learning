package main

import (
	"fmt"
)

// Finding a remainder = Calculating "a modulus"

func main() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue // skip whenever there is an even number, thus getting all odd numbers...
		}
		fmt.Printf("This is an odd number \t %v\n", i)
	}
}

/* continue
// takes to next iteration

for i := 0; i < 20; i++ {
	if i%2 != 0 {
		continue
	}
	fmt.Println("Counting even numbers: ", i)
}
*/
