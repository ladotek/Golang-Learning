package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c))) // Polymorphism - Value of one type can also be a value of another type
}

func main() {
	b := book{
		title: "West with the Night",
	}

	var n count = 42

	fmt.Println(b)
	fmt.Println(n)

}
