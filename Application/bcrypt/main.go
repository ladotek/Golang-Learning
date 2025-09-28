package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1)) // Compares "bs" - which is the hashed password, and the loginPword1, throws error if doesnt match...
	if err != nil {
		fmt.Println("WRONG PASSWORD")
		return
	}

	fmt.Println("Sucessful login")

}

/*

A hashing algorithm implementation

*/
