package main

import (
	"fmt"
	"time"
)

// Wrapper function for adding timing information
func TimedFunction(fn func()) { // Function "TimedFunction" takes in the function "fn"
	start := time.Now()          // The "time" package (that's imported) gives the current time
	fn()                         // Then we run the function that's passed in
	elapsed := time.Since(start) // Since function, will return a duration
	fmt.Println("Elapsed time:", elapsed)
}

// WE ARE WRAPPING THE fn() function with some processing, with something before it and something after it!!

// Function to be wrapped
func MyFunction() {
	time.Sleep(2 * time.Second) // Simulate some work
	fmt.Println("MyFunction completed")
}
func main() {
	// Call the wrapped function
	TimedFunction(MyFunction)
}

/*
To summarize, a wrapper function wraps or modifies another function's behavior, while a
callback function is a function passed as an argument to be executed at a specific point or
event.

GOOD VIDEO TO WATCH ON UDEMY TO LEARN!!!!!
*/
