package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Guinness", "Wolverine Tracks"}
	fmt.Println(jb)
	fmt.Println(jm)

	//A slice holding slices...First two brackets make it a slice, and its storing slices of type string...
	xxs := [][]string{jb, jm}
	fmt.Println(xxs)
}
