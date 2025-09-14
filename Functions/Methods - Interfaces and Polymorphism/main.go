package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() { //This defines a method named speak on the person type. \\
	fmt.Println("I am", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface { // the human interface has 1 method - speak
	speak() // any type that has this method (types: secretAgent and person), is also of this type (human)!
}

func saySomething(h human) {
	h.speak()
}

/*
(sa person) is the receiver — it means this method is "attached" to person
and can be called with sa.speak() or p2.speak().
*/
func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	p2 := person{
		first: "Jenny",
	}

	/*sa1.speak()
	p2.speak()
	*/
	saySomething(sa1)
	saySomething(p2)
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }
