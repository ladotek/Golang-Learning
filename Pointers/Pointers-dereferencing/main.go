package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x) // & gives the memory address of where the value is stored (0x...)

	y := &x // y and &x are now both pointing to the same memory address

	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)  // Dereferencing y, the asterisk (*) gives the value of the memory address
	fmt.Println(*&x) // Dereferencing

	*y = 43 // Make the value of that memory address = 43 --- ALso changes the memory address value of x to = 43
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*y)

}

// y and &x are both pointing to the same memory address, so when we assign 43 as the value of memory address *y, it changes the value of x
