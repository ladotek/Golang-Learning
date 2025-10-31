package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100        // gs = 100 goroutines to launch.
	var wg sync.WaitGroup // Create a WaitGroup so main() can wait for them all to finish.
	wg.Add(gs)            // wg.Add(gs) says: “I’ll wait for 100 goroutines to call Done().”

	for i := 0; i < gs; i++ { // Loops 100 times → spawns 100 new goroutines.
		go func() { // Each iteration runs concurrently, not sequentially.
			v := counter // Reads the current value of counter into a local variable v.
			// time.Sleep(time.Second)
			runtime.Gosched() // Yields the processor, telling the Go scheduler: “Let other goroutines run now.”This increases the chance of a race condition (since multiple goroutines will now access counter at the same time).
			v++
			counter = v // Increments the local copy, then writes back to counter. ...Because other goroutines are doing the same thing at the same time, their reads/writes interfere. → Race condition: some increments get lost.
			wg.Done()   //Decrements the WaitGroup count by 1 (signals that this goroutine is done).
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine()) // Prints how many goroutines currently exist.

	}
	wg.Wait()                                          // Blocks main() until the WaitGroup counter reaches 0 (i.e., all 100 goroutines finished).
	fmt.Println("Goroutines:", runtime.NumGoroutine()) // After wg.Wait(), usually back to 1 goroutine (main).
	fmt.Println("count:", counter)                     // Prints final counter value.
}

/*
What happens when you run it

Expected theoretical result: counter = 100
Typical actual result: anywhere from 1 to 100 (often much lower).

Because all goroutines:

read the same value of counter,

yield,

then overwrite each other’s increments.

This is a data race.

*/
