package main

import "fmt"

func main() {

	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) // signal: "no more values will be sent" ----This is essential. If we dont close the channel, it goes into DEADLOCK, waiting for more values to come, but cant
	}() // End of the sending goroutine. ---- () after } calls this function immediately after defining it...Calls it concurrently

	// receive
	for v := range c { // range c repeatedly does v, ok := <-c; when close(c) has happened and all values are consumed, ok becomes false and the loop exits.
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
