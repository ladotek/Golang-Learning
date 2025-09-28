package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Fprintln(os.Stdout, "Hello, playground") // This is exactly what "Println" does with the help of interfaces...!

	io.WriteString(os.Stdout, "Hello, playground")
}

/*

https://pkg.go.dev/io#WriteString


*/
