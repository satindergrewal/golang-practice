// basic-middleware.go
package main

import (
	"fmt"
	"log"
	"time"
	"math/rand"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		if r.URL.Path == `/foo` {
			fmt.Println("foo was accessed")
		}
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
	go func(){
		//time.Sleep(2000 * time.Millisecond)
		log.Println("hello from foo handler after 5 seconds of request recieving")
		
		out := make(chan string)
		go func(){
			for i := 0; ; i++ {
				out <- `foo said something`
				time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			}
			//out <- `foo said something`
		}()
		//go receiver(out)
		printer(out)
	}()
}

func receiver(out chan<- string) {
	//out <- `some update`
	go func() {
		for {
			log.Println(out)
		}
	}()
}

func printer(in <-chan string) {
	//log.Println(<-in)
	go func() {
		for {
			log.Println(<-in)
		}
	}()
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)
}
