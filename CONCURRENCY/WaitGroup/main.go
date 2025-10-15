package main

import (
	"fmt"
	"runtime" //info about the Go runtime/OS/CPUs/goroutines.
	"sync"    //concurrency primitives (here: WaitGroup).
)

var wg sync.WaitGroup // A package-level WaitGroup used to wait for goroutines to finish.

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("Architecture\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOroutines\t", runtime.NumGoroutine())

	wg.Add(1) // Increments the WaitGroup counter to 1, signaling that one goroutine’s completion will be awaited.(Call Add before starting the goroutine to avoid races.)

	go foo() // By adding "go" in front of the function, it now runs on concurrent design pattern and launches a new goroutine!
	//Starts foo() concurrently. Now there are (at least) 2 goroutines: main and foo.
	bar() // Calls bar() synchronously in the main goroutine. While bar runs, foo runs in parallel (subject to scheduler/CPUs).

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GOroutines\t", runtime.NumGoroutine())
	wg.Wait() // Blocks the main goroutine until the WaitGroup counter returns to 0 (i.e., until foo calls wg.Done()).
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done() // Calls wg.Done() to decrement the WaitGroup counter (from 1 → 0).

}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}

/*
Why Wait is needed: Without wg.Wait(), main could exit before foo finishes, terminating the process early.

Goroutine count: First print (before go foo()) ≈ 1. Second print (after bar()) is nondeterministic (often 1 or 2).
*/
