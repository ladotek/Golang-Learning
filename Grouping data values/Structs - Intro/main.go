package main

import "fmt"

/*
A struct is an composite data type. (composite data types, aka, aggregate data types, aka,
complex data types). Structs allow us to compose together values of different types.
*/

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", p1, p1)
	fmt.Println(p1.first, p1.last, p1.age)

}
