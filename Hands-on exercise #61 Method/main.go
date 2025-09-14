package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() { //This defines a method named speak on the person type. \\
	fmt.Println("I am", p.first, "and I am ", p.age, "years old")
}

func main() {
	p1 := person{
		first: "James",
		age:   42,
	}

	p1.speak()

}

/*
Hands-on exercise
● Create a user defined struct with
	○ the identifier “person”
	○ the fields:
		■ first
		■ age
● attach a method to type person with
	○ the identifier “speak”
	○ the method should have the person say their name and age
● create a value of type person
● call the method from the value of type person

*/
