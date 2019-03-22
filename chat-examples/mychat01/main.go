package main

import (
	"fmt"
)

type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handler() {
	ch := make(chan string) // outgoing client messages
	for {
		go fmt.Println(ch)
	}

	who := "John"
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch
}

func main() {
	fmt.Println("hello")

	go broadcaster()
	for {
		go handler()
	}
}