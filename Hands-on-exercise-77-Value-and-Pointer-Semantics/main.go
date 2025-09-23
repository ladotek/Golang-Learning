package main

import "fmt"

type person struct {
	first string
}

func main() {
	p := person{
		first: "Jenny",
	}
	fmt.Println(p)
	p = changeName(p, "Jen")
	fmt.Println(p)

	changeNamePointers(&p, "Jennifer")
	fmt.Println(p)

}

func changeName(p person, s string) person { // Using value semantics!
	p.first = s
	return p
}

func changeNamePointers(p *person, s string) { // Using pointer semantics!
	p.first = s
}

/*

Create two functions to change a field in a struct called "first" of TYPE string:
● One function will use VALUE SEMANTICS
	○ this function will return some VALUE of some TYPE
● The other function will use POINTER SEMANTICS
	○ this function will return nothing
● don't use methods

*/
