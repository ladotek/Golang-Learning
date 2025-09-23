package main

import "fmt"

func main() {
	a := 1
	fmt.Println(&a)
}

/*
go run -gcflags -m main.go

escapes to the heap: variable shared between main() and Println() --	fmt.Println(a)


moved to heap: variable moved to the heap -- 	fmt.Println(&a)

*/
