package main

import "fmt"

func main() {
	/* WRONG!:
	an := map[string][]string{
		"bond_james":       {"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": {`intelligence`, `literature`, `computer science`},
		"no_dr":            {`cats`, `ice cream`, `sunsets`},
	}
	for i, v := range an {
		fmt.Printf("Key: %v - Value: %v\n", i, v)
	}

	*/

	m := make(map[string][]string) // key = [string] | value = []string // The key is a string, the value is a slice of type string

	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	fmt.Printf("%#v\n", m)

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v { //Nested Loop!
			fmt.Println(i, v2)
		}
	}

}

/*
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of
TYPE []string which stores the
ir favorite things. Store three records in your map. Print out all of
the values, along with their index position in the slice.
key value
`bond_james` `shaken, not stirred`, `martinis`, `fast cars`
`moneypenny_jenny` `intelligence`, `literature`, `computer science`
`no_dr` `cats`, `ice cream`, `sunsets`

*/
