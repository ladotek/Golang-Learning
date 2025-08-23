package main

import (
	"fmt"
	"log"
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

func logInfo(s fmt.Stringer) { // Wrapper function
	log.Println("LOG FROM 138", s.String()) // This function is being wrapped by the logInfo function and adding a functionlity - "LOG FROM"
}

func main() {
	b := book{
		title: "West with the Night",
	}

	var n count = 42

	logInfo(b)
	logInfo(n)

}
