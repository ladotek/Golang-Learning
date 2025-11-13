package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 42

	fmt.Println(<-c)

}

/*

when you send and receive on a channel, the transaction can not occur until send and receive can happen at the same time. If they cant, its blocked!

CHANNELS BLOCK!

*/
