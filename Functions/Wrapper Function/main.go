package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := ReadFile("poem.txt")
	if err != nil {
		log.Fatalf("Error in main in readFile: %s", err)
	}
	fmt.Println(xb) // Prints the byte slice itself first. Basically prints a []byte (aka []uint8) in Go’s default format: the decimal byte values, e.g.
	fmt.Println(string(xb))

}

func ReadFile(fileName string) ([]byte, error) { // Takes in a file name (that is a string), and returns a slice of bytes and an error(if there is one)
	xb, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("There was an error in readFile: %s", err)
	}
	return xb, nil
}

// A named function:
//// func (r receiver) identifier(p parameter(s)) (return(s)) { <your code here> }

//An anonymous function
// func(p parameter(s)) (r return() { code }
