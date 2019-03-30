package main

import (
	"flag"
	"fmt"
	"goblueprints/chapter1/trace"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/gorilla/websocket"
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

	// Creating the URL scheme to use with websocket Dialer
	// to connnect to ws://localhost:8080/room
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/room"}
	log.Printf("connecting to %s", u.String())

	// Initiate the websocket connection from the go code **as a client** to connect to the chat room
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(8e3)) * time.Millisecond)
			// Just printing the log of the same message in command line. Might better to ignore it.
			// log.Println("Sending automatic hello from root ServeHTTP handle to web page!")

			// Write the Message as Text message to the web socket connection
			// which will show up in the chat box
			err := c.WriteMessage(websocket.TextMessage, []byte("Sending automatic hello from root ServeHTTP handle to web page!"))
			if err != nil {
				log.Println("write:", err)
				return
			}
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
