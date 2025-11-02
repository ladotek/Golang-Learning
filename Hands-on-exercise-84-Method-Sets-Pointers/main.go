package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() { //This defines a method named speak on the person type. \\
	fmt.Println("I am", p.first)
}

func saySomething(h human) {
	h.speak() // !!!! 	○ have it take in a human as a parameter --- have it call the speak method
}

func main() {

	p1 := person{
		first: "Lado",
	}

	saySomething(&p1)
	//saySometing(p1)  ---> This will not work!!

	p1.speak() // This will work no matter what

}

/*
	This exercise will reinforce our understanding of method sets:
● create a type person struct
● attach a method speak to type person using a pointer receiver
	○ *person
● create a type human interface
	○ to implicitly implement the interface, a human must have the speak method
● create func “saySomething”
	○ have it take in a human as a parameter
	○ have it call the speak method
● show the following in your code
	○ you CAN pass a value of type *person into saySomething
	○ you CANNOT pass a value of type person into saySomething
● here is a hint if you need some help

Receivers Values
-----------------------------------------------
(t T) T and *T
(t *T) *T


*/
