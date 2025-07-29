package main

import (
	"fmt"
)

func main() {
	x := 20

	for {
		if x < 0 {
			break
		}
		fmt.Println(x)
		x--
	}

}

/* INFINITE LOOP! :) Will keep printing numbers below 20 infinitely

func main() {
	x := 20

	for {
		fmt.Println(x)
		x--
	}

}
*/
