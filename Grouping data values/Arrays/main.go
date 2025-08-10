package main

import "fmt"

func main() {
	//array literal
	a := [3]int{42, 43, 44}
	fmt.Println(a)

	b := [...]string{"Hello", "Gophers!"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7 // Arrays and slices use zero based indexes! The "first" position (number) is 0 !
	c[1] = 8

	fmt.Println(c)

	fmt.Printf("%T\t %T\t %T\n", a, b, c)

	{
		//declare a variable of type [7]int
		var ni [7]int
		//assign a value to index position zero
		ni[0] = 42
		fmt.Printf("%#v \t\t\t\t %T\n", ni, ni)

		//array literal
		ni2 := [4]int{55, 56, 57, 58}
		fmt.Printf("%#v \t\t\t\t\t %T\n", ni2, ni2)

		//array literal
		//have compiler count elements
		ns := [...]string{"Chocolate", "Vanilla", "Strawberry"}
		fmt.Printf("%#v \t %T\n", ns, ns)

		//use the built-in len function
		// https://pkg.go.dev/builtin#len
		fmt.Println(len(ni))
		fmt.Println(len(ni2))
		fmt.Println(len(ns))

	}
}
