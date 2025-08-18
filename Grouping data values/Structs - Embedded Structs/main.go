package main

import "fmt"

/*
We can take one struct and embed it in another struct. When you do this the inner type gets
promoted to the outer type.
*/

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk   bool
	first string
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "ETHAN HAWKE", // When printing sa(secretAgent), it will print ETHAN HAWKE, because its associated with it directly!
		ltk:   true,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1)
	fmt.Println(p2)

	fmt.Printf("%T \t %#v\n", sa1, sa1)

	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.person)
	fmt.Println(sa1.first)
	fmt.Println(sa1.person.first)

}
