package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr. No"}

	fmt.Println("Unsorted: ", xi)
	fmt.Println("Unsorted: ", xs)

	fmt.Println("----------------------")

	sort.Ints(xi)
	fmt.Println("Sorted: ", xi)

	sort.Strings(xs) // This sorts alphabetically!
	fmt.Println("Sorted: ", xs)

}

/*
https://pkg.go.dev/sort
*/
