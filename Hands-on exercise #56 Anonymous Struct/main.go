package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Bob":   5,
			"Carol": 3,
		},
		favDrinks: []string{"Beer", "Coke", "Champagne"},
	}

	/* Print the whole struct
	fmt.Printf("Full struct: %#v\n\n", p1)

	// Print individual fields
	fmt.Println("First name:", p1.first)

	fmt.Println("\nFriends and friendship level:")
	for name, level := range p1.friends {
		fmt.Printf("- %s: %d\n", name, level)
	}

	fmt.Println("\nFavorite drinks:")
	for i, drink := range p1.favDrinks {
		fmt.Printf("%d. %s\n", i+1, drink)
	}
	*/
	for k, v := range p1.friends {
		fmt.Println(p1.first, "- friends -", k, v)
	}

	for k, v := range p1.favDrinks {
		fmt.Println(p1.first, "- drinks -", k, v)
	}
}

/*
Create and use an anonymous struct with these fields:
	● first string
	● friends map[string]int
	● favDrinks []string
Print things.
*/
