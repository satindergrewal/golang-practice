package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// gospot.mchampaneri.in ===> localhost:8085
	u1, _ := url.Parse("http://localhost:8080/")
	http.Handle("app2.domain.com/", httputil.NewSingleHostReverseProxy(u1))

	// www.mchampaneri.in ===> localhost:8081
	u2, _ := url.Parse("http://localhost:8081/")
	http.Handle("www.mchampaneri.in/", httputil.NewSingleHostReverseProxy(u2))

	// Start the server
	http.ListenAndServe(":80", nil)
}
