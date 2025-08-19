package main

import "fmt"

type person struct {
	first    string
	last     string
	icecream []string
}

func main() {
	p1 := person{
		first:    "James",
		last:     "Bond",
		icecream: []string{"Chocolate", "Vanilla", "Strawberry"},
	}

	p2 := person{
		first:    "Jenny",
		last:     "Moneypenny",
		icecream: []string{"Cookie", "Kiwi", "Raspberry"},
	}

	fmt.Println(p1, p2)

	for _, v := range p1.icecream {
		fmt.Println(p1.first, "favorite is", v)
	}

	for _, v := range p2.icecream {
		fmt.Println(p2.first, "favorite is", v)
	}
}

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
● first name
● last name
● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.
*/
