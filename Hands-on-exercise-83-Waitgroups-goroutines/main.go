package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin CPU", runtime.NumCPU())
	fmt.Println("begin goroutines", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from thing two")
		wg.Done()
	}()

	fmt.Println("mid CPU", runtime.NumCPU())
	fmt.Println("mid goroutines", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("About to exit")

	fmt.Println("end CPU", runtime.NumCPU())
	fmt.Println("end goroutines", runtime.NumGoroutine())
}

/*
	● in addition to the main goroutine, launch two additional goroutines
		○ each additional goroutine should print something out
	● use waitgroups to make sure each goroutine finishes before your program exists
*/
