package main

import "fmt"

func main() {
	foo()

	bar("Todd")

	s := aloha("Todd")
	fmt.Println(s)

	s1, n := dogYears("Todd", 40)
	fmt.Println(s1, n)
}

func foo() { //no params, no returns
	fmt.Println("I am from foo")
}

func bar(s string) { //1 param, no returns
	fmt.Println("My name is", s)
}

func aloha(s string) string { //1 param, 1 return
	return fmt.Sprint("They call me Mr ", s)
}

func dogYears(name string, age int) (string, int) { // 2 params, 2 returns
	age *= 7
	return fmt.Sprint(name, " is this old in dog years "), age
}

// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }
