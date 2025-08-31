package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct { // A simple struct holding a first name
	first string
}

func (p person) writeOut(w io.Writer) { // Method with interface param
	w.Write([]byte(p.first)) // Sends the bytes of the name to whatever writer you pass (file, buffer, network, stdout, etc.).
}

func main() {
	p := person{ //Defines a person with a first name.
		first: "Jenny",
	}

	f, err := os.Create("output.txt") // Creates/truncates a file and returns an *os.File, which implements io.Writer.
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close() // ensures the file is closed when main returns (even on error)

	var b bytes.Buffer // Important: bytes.Buffer’s Write method has a pointer receiver, so only *bytes.Buffer satisfies io.Writer. That’s why the code passes &b to writeOut.
	p.writeOut(f)
	p.writeOut(&b)
	fmt.Println(b.String())
}

/*
Execution flow

- Construct p := person{first: "Jenny"}.
- Create output.txt → f.
- Create an empty buffer → b.
- p.writeOut(f) writes "Jenny" into the file.
- p.writeOut(&b) writes "Jenny" into the buffer.
- fmt.Println(b.String()) prints Jenny to stdout (with a newline from Println).

Result:
- The file output.txt contains: Jenny
- The terminal prints: Jenny

*/
