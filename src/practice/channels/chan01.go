package main

import (
	"fmt"
)

// receiver function with value parameter of only reciving channel variable of type integer
func receiver(out chan<- int) {
	// Stores just single integer value to it's channel variable
	out <- 20
}

// printer function with value parameter of only sending channel variable of type integer
func printer(in <-chan int) {
	// Prints the value from the variable in
	fmt.Println(<-in)
}

func main() {
	fmt.Println("Hello Channels with functions!")

	// Declaration of integer channel in main function
	// Step 1 - Create a Unidirectional channel of type integer
	out := make(chan int)

	// Goroutine 2
	// Step 2 - Execute the reciever function as concurently goroutine with the channel created in main function body.
	// The out channel in main body gets the integer value 20 stored from reciever function
	go receiver(out)

	// main Goroutine/Goroutine 1
	// Step 3 - Since other functions in main body executing concurently printer function will execute concurently being in a main body's concurent process.
	// From the earlier concurent process out channel declared in main body holds the integer value of 20 from reciever function
	// passing the same main body channel out to printer sends this value to printer function
	// printer function just prints that channel value in it's printer function body
	printer(out)
}