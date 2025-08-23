package main

import "fmt"

type person struct {
	first string
}

func (p person) speak() { //This defines a method named speak on the person type. \\
	fmt.Println("I am", p.first)
}

/*
(p person) is the receiver — it means this method is "attached" to person
and can be called with p1.speak() or p2.speak().
*/
func main() {
	p1 := person{ //Two person structs are created and initialized with different first names.
		first: "James",
	}
	p2 := person{
		first: "Jenny",
	}

	p1.speak() //Each person instance calls its own .speak() method.
	p2.speak()
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }
