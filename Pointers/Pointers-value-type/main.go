package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x) // &x gives the address of where the value is stored (0x...)

	fmt.Printf("%v\t%T\n", &x, &x)

	s := "James"

	fmt.Printf("%v\t%T\n", &s, &s)
}
