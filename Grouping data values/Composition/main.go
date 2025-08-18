package main

import "fmt"

type Engine struct {
	// Engine fields
}

// engine method
func (e *Engine) Start() {
	fmt.Println("Engine started")
}

type Car struct {
	Engine // Embedding the Engine struct
	// Car-specific fields
}

func main() {
	car := Car{}
	car.Start() // Calling the Start() method from the embedded Engine struct
}
