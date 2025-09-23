package main

import "fmt"

// Value semantics --> arguments are copied. The callee cannot change the caller’s variable.
// Defines addOne, which takes a copy of an int (v) and returns v + 1.
// Because v is a copy, the caller’s variable is not modified.
func addOne(v int) int {
	return v + 1
}

// Pointer semantics --> you pass an address; the callee can dereference and mutate the original.
// Defines addOneP, which takes a pointer to int (*int).
// *v dereferences the pointer (accesses the original int), then += 1 modifies the original in place.
func addOneP(v *int) {
	*v += 1
}

func main() {
	//Value semantics
	a := 1
	fmt.Println(a)         //1
	fmt.Println(addOne(a)) //2 -- Calls addOne(a). "a" is passed by value (a copy), so addOne returns 2.

	fmt.Println(a) //1 -- Prints "a" again; still 1 because addOne changed only its local copy.

	//Pointer semantics
	b := 1
	fmt.Println("value semantics")
	fmt.Println(b)
	addOneP(&b)    //&b takes the address of b (pointer to b) and passes it to addOneP. Inside addOneP, *v += 1 mutates the original b.
	fmt.Println(b) //2 -- Prints "b" again; now it’s 2 because the pointer-based function (addOneP) modified b in place.
}
