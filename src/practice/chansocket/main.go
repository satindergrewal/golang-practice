package main

import (
	"io"
	"net/http"
	"fmt"
	"github.com/gorilla/websocket"
)

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

/*
var ws = new WebSocket("ws://localhost:8080/ws")
ws.addEventListener("message", function(e) {console.log(e);});
ws.onmessage = function (event) {
	console.log(event.data);
}

ws.readyState
ws.CLOSED
ws.OPEN
ws.close()

ws.send("foo")
ws.send(JSON.stringify({username: "Sat"}))
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
	http.HandleFunc("/ws", ws)

	http.ListenAndServe(":8080", nil)
}


