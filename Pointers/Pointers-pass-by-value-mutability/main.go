package main

import "fmt"

func intDelta(n *int) { //intDelta takes a pointer to int.
	*n = 43 // *n = 43 dereferences the pointer and sets the original int to 43.
}

func sliceDelta(ii []int) { //sliceDelta takes a slice of ints.
	ii[0] = 99 // Puts 99 at beginning of slice of int...
}

func mapDelta(md map[string]int, s string) { //mapDelta takes a map and a key.
	md[s] = 33 //Maps are also descriptors pointing to a runtime hash table. Mutations (sets/deletes) affect the shared table, so the caller sees the change.
}

func main() {
	a := 42
	fmt.Println(a)

	intDelta(&a)   //Pass address of a to intDelta. Inside, *n = 43 changes the original.
	fmt.Println(a) // This will now print 43.

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	sliceDelta(xi)  //Pass the slice descriptor. sliceDelta sets index 0 to 99 in the same underlying array.
	fmt.Println(xi) // Prints [99 2 3 4].

	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"])
	mapDelta(m, "James")    //Pass the map descriptor and key. mapDelta writes 33 under "James" to the same underlying map.
	fmt.Println(m["James"]) //Prints 33.
}
