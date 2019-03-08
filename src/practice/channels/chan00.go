package main

import (
	"fmt"
)

var SayHello string = "Hello from Init!"

func main() {
	fmt.Println(SayHello)
	fmt.Println("Hello Channels!")

	out := make(chan int)
	in := make(chan int)

	// Goroutine 1
	// out<- receving value and storing it in out channel
	// Value passed as integer value.
	go func() {
		out<- 10
	}()

	// Goroutine 2
	// in<- is receving value and storing it in in channel. 
	// Value passed by the other channel "out", which holds the integer value, and it is "sent" using <-out
	// Running "in" concurantly to "out" anonymous go func so it can get the value from out and store to it.
	go func() {
		in<- <-out
	}()

	// main Goroutine, counts as the mother go routine
	fmt.Println(<-in)
}