package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup // A WaitGroup to wait for goroutines to finish. Zero value is ready to use. Methods have pointer receivers, but you can call them on a value; the compiler handles the address automatically.
	var incrementer int64 // int64 usually indicates that the package atomic is being used

	gs := 100  // Number of goroutines (goroutine span).
	wg.Add(gs) // Tells the WaitGroup to expect gs calls to Done(). Must be called before starting the goroutines.

	for i := 0; i < gs; i++ { // Loop that launches 100 goroutines. (Note: i isnt used.)
		go func() { // Start a new goroutine running this anonymous function. It captures (closes over) incrementer and wg from the surrounding scope.
			atomic.AddInt64(&incrementer, 1)
			fmt.Println(atomic.LoadInt64(&incrementer))
			wg.Done() // Signal that this goroutine finished one unit of work (decrements the WaitGroup counter by 1).
		}() // Invoke the anonymous function immediately (still running concurrently due to go).
	}
	wg.Wait()
	fmt.Println("end value", incrementer)
}

/*
Fix the race condition you created in previous exercise by using package atomic

*/
