package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup // A WaitGroup to wait for goroutines to finish. Zero value is ready to use. Methods have pointer receivers, but you can call them on a value; the compiler handles the address automatically.

	incrementer := 0 // Shared counter that all goroutines will read and write (this is the source of the race).
	gs := 100        // Number of goroutines (goroutine span).
	wg.Add(gs)       // Tells the WaitGroup to expect gs calls to Done(). Must be called before starting the goroutines.

	for i := 0; i < gs; i++ { // Loop that launches 100 goroutines. (Note: i isnt used.)
		go func() { // Start a new goroutine running this anonymous function. It captures (closes over) incrementer and wg from the surrounding scope.
			v := incrementer         // Read the shared variable into a local copy v. This read races with other goroutines’ writes.
			runtime.Gosched()        // Yield the processor so other goroutines can run now. This increases the chance that multiple goroutines interleave between read and write, exposing the race. It is not a synchronization mechanism.
			v++                      // Increment the local copy.
			incrementer = v          // Write back to the shared variable. This write races with other goroutines’ reads/writes. Classic read-modify-write race.
			fmt.Println(incrementer) // Print the (currently) stored value. Interleaved prints are expected and also race with other writes.
			wg.Done()                // Signal that this goroutine finished one unit of work (decrements the WaitGroup counter by 1).
		}() // Invoke the anonymous function immediately (still running concurrently due to go).

	}
	wg.Wait()                             // Block the main goroutine until all 100 goroutines have called Done().
	fmt.Println("end value", incrementer) // Print the final value after all goroutines are done. Even though the read here is after all writes, many updates were lost earlier due to the race, so this can be anything from ~1 to ~100.
}

/*
 Using goroutines, create an incrementer program
	○ have a variable to hold the incrementer value
	○ launch a bunch of goroutines
		■ each goroutine should
			● read the incrementer value
				○ store it in a new variable
			● yield the processor with runtime.Gosched()
			● increment the new variable
			● write the value in the new variable back to the incrementer variable
● use waitgroups to wait for all of your goroutines to finish
● the above will create a race condition.
● Prove that it is a race condition by using the -race flag



*/
