// https://github.com/fasthttp/router/tree/master/_examples/basic
package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// Index is the index handler
func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

// Hello is the Hello handler
func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

// MultiParams is the multi params handler
func MultiParams(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hi, %s, %s!\n", ctx.UserValue("name"), ctx.UserValue("word"))
}

// QueryArgs is used for uri query args test #11:
// if the req uri is /ping?name=foo, output: Pong! foo
// if the req uri is /piNg?name=foo, redirect to /ping, output: Pong!
func QueryArgs(ctx *fasthttp.RequestCtx) {
	name := ctx.QueryArgs().Peek("name")
	fmt.Fprintf(ctx, "Pong! %s\n", string(name))
}

func main() {
	r := router.New()
	r.GET("/", Index)
	r.GET("/hello/{name}", Hello)
	r.GET("/multi/{name}/{word}", MultiParams)
	r.GET("/ping", QueryArgs)

	log.Fatal(fasthttp.ListenAndServe(":8080", r.Handler))
}

// ➜  ~ curl http://localhost:8080/
// Welcome!
// ➜  ~ curl http://localhost:8080/hello/satinder
// hello, satinder!
// ➜  ~ curl http://localhost:8080/multi/satinder/grewal
// hi, satinder, grewal!
// ➜  ~
