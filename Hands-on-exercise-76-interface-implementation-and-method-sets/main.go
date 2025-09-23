package main

// This exercise was solved using the "Use Value Semantics When Possible" suggestion! So we just changed the codebase accordingly...

import "fmt"

type dog struct { // Define a struct type dog with one field, first (the dog’s name).
	first string
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I'm running")
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Henry"}
	youngRun(d1)

	d2 := dog{"Padget"}
	youngRun(d2)

	d2 = d2.changeName("Rover")
	youngRun(d2)
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}

/*

In the provided code, correct the code to do the following:
● implement the "youngin" interface
	○ you can change anything in the code to do so except the interface
● maintain consistency within the code
	○ receivers should stick with either POINTER or VALUE semantics
● when choosing between POINTER or VALUE semantics, understand why you chose one
or the other

*/
