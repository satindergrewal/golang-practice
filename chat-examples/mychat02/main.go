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
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			fmt.Println("Message to console from root ServeHTTP handle")
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
