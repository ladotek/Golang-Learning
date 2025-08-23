package main

import "fmt"

/*
We can create anonymous structs also. An anonymous struct is a struct which is not
associated with a specific identifier.
*/

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{first: "James",
		last: "Bond",
		age:  32,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Printf("%T\t , %v \n", p1, p1)
	fmt.Printf("%T\t , %v \n", p2, p2)

}
