package main

import (
	"flag"
	"fmt"
	"goblueprints/chapter1/trace"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	fmt.Println(r.Host)
	t.templ.Execute(w, r)

	room := newRoom()
	socket, _ := upgrader.Upgrade(w, r, nil)
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   room,
	}
	room.join <- client
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(8e3)) * time.Millisecond)
			room.forward <- []byte("Hello from / ServeHTTP Handle")
			//client.socket.WriteMessage(websocket.TextMessage, []byte("Hello from / ServeHTTP Handle"))
			//fmt.Println("Sending automatic hello from root ServeHTTP handle to web page!")
		}
	}()
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// get the room going
	go r.run()

	//  Start the Web Server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
