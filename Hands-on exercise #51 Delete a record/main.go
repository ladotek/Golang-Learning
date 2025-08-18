package main

import "fmt"

func main() {
	m := make(map[string][]string) // key = [string] | value = []string // The key is a string, the value is a slice of type string

	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}

}

/*
Using the code from the previous example, delete a record from your map. Now print the map
out using the “range” loop
*/
