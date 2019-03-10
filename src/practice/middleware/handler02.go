// basic-middleware.go
package main

import (
	"fmt"
	"log"
	"time"
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
		time.Sleep(5000 * time.Millisecond)
			fmt.Println("hello from foo handler after 5 seconds of request recieving")
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
