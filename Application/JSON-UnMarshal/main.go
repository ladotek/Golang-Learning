package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First`
	Last  string `json:"Last`
	Age   int    `json:"Age`
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]` // That's our JSON -- DONT FORGET THE `` surrounding the JSON!!!!!
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{} --- CAN ALSO USE THIS AND WILL WORK!
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nall of the data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}

/* Output:

string
[]uint8   <---- uint8 is an alias for byte!


OUTPUT:
all of the data [{James Bond 32} {Miss Moneypenny 27}]

PERSON NUMBER 0
James Bond 32

PERSON NUMBER 1
Miss Moneypenny 27



*/
