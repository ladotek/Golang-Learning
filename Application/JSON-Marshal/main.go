package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}

/*
Output:
[{James Bond 32} {Miss Moneypenny 27}] <--- This is our code and it is changed to JSON below:

[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]  <--- JSON! Can be sent
*/
