package main

import (
	"io"
	"net/http"
	"fmt"
	"math/rand"
	"log"
	"time"
	"github.com/gorilla/websocket"
)


func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
	
	//Some code here who's output I want to pass to websockets url ws://localhost:8080/ws
	n := timeConsumingWork(4)
	fmt.Println("Random Number Print from cat: ", n)
	//Example the value of n I need to pass to ws://localhost:8080/ws, how can I do it?

	// Some other example test code just giving some random output from hotcat http handler
	// Would like to pass it's output to ws://localhost:8080/ws to print in websocckets output in browser
	go func(){
		out := make(chan string)
		go func(){
			for i := 0; ; i++ {
				out <- `foo said something`
				time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			}
			//out <- `foo said something`
		}()
		printer(out)
	}()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Execute this in browser console to initiate websococket connection and to send ws.send() commands etc.
/*
var ws = new WebSocket("ws://localhost:8080/ws")
ws.addEventListener("message", function(e) {console.log(e);});
ws.onmessage = function (event) {
	console.log(event.data);
}

ws.send("foo")
ws.send(JSON.stringify({username: "Sat"}))

ws.readyState
ws.CLOSED
ws.OPEN
ws.close()
*/
func ws(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msgType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msgType, msg); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	var c hotcat

	http.Handle("/cat", c)
	http.HandleFunc("/ws", logging(ws))

	http.ListenAndServe(":8080", nil)
}


func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}


func printer(in <-chan string) {
	//log.Println(<-in)
	go func() {
		for {
			log.Println(<-in)
		}
	}()
}