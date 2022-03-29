package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

//Target url: https://httpbin.org/headers
//Url through proxy:  http://localhost:3002/forward/headers

func main() {
	target := "https://httpbin.org"
	remote, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	r := mux.NewRouter()
	r.HandleFunc("/forward/{rest:.*}", handler(proxy))
	http.Handle("/", r)
	http.ListenAndServe(":3002", r)
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = mux.Vars(r)["rest"]
		p.ServeHTTP(w, r)
	}
}
