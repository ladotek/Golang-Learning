package main

import "fmt"

var y int

func main() {
	fmt.Println(y) //var for zero value

	z := 42 //short declaration operator
	fmt.Println(z)

	a, b := 43, "Happiness" //multiple initializations
	println(a, b)

	var c float32 = 42.42 //var when specificity is required
	fmt.Printf("%v is of this type %T\n", c, c)

	e, f, _ := 45, 46, 47 //blank identifier
	fmt.Println(e, f)

}

/*
var for zero value
short declaration operator
multiple initializations
var when specificity is required
blank identifier

*/
