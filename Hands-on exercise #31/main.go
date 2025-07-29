package main

import (
	"fmt"
)

func main() {
	i := 20
	for i >= -5 {
		fmt.Println(i)
		i-- //Decrementer, going into negative...
	}
}
