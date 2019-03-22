//authserver.go
package main

import(
    "fmt"
    "net/http"
    "github.com/goji/httpauth"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "some message to authenticated user.\n")
}

func main() {
    http.Handle("/", httpauth.SimpleBasicAuth("someuser", "somepassword")(http.HandlerFunc(YourHandler)))
    http.ListenAndServe(":7000", nil)
}