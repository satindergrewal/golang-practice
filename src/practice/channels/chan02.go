package main

import (
	"fmt"
	"time"
)
func receiver(out chan<- int) {
	out <- 20
}
func printer(in <-chan int) {
	fmt.Println(<-in)
}

func saySomthing() {
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("Saying something after 5 seconds of delay")
}

func main() {
	fmt.Println("Hello Channels with functions!")

	out := make(chan int)

	go receiver(out)

	printer(out)

	saySomthing()
}