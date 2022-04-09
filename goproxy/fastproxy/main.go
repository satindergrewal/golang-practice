package main

import (
	"log"
	"strings"
	"time"

	// "github.com/rs/zerolog/log"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

// Some config consts
const (
	// Host where proxy will redirect requests
	ApiHost = "localhost:8080"
	// Clients timeout
	TimeOut = time.Second * 10
	// Port where proxy start
	ProxyPort = ":80"
	// Name of proxy server
	ProxyName = "proxy"
	// Prefix of path for redirect
	ApiPath = "/"
)

func main() {
	// log.Info().Msg(ProxyName + " start")
	log.Println(ProxyName + " start")

	// handler for all requests coming to proxy
	proxyHandler := func(ctx *fasthttp.RequestCtx) {
		switch {
		case strings.HasPrefix(string(ctx.Path()), ApiPath): // if path has prefix, its proxying to api host
			ctx.Request.SetHost(ApiHost)                                    // just change host of coming request
			err := fasthttp.DoTimeout(&ctx.Request, &ctx.Response, TimeOut) // and do request
			if err != nil {
				// log.Error().Err(err).Send() // log err
				log.Println(err)
				return
			}
		}
	}

	r := router.New()
	r.GET("/", proxyHandler)

	// server setup
	s := &fasthttp.Server{
		Handler:     proxyHandler,
		Name:        ProxyName,
		ReadTimeout: TimeOut * 2,
	}

	// log.Info().Msg(ProxyName + " ok")
	log.Println(ProxyName + " ok")

	// fasthttp.ListenAndServe(":"+fmt.Sprintf("%d", 3334), r.Handler)

	if err := s.ListenAndServe(ProxyPort); err != nil { // start listening. ListenAndServe() block func
		// log.Fatal().Err(err).Send()
		log.Println(err)
	}
}
