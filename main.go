package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	server := fasthttp.Server{
		Handler: requestHandler,
	}

	if err := server.ListenAndServe(":8080"); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/hello":
		helloHandler(ctx)
	default:
		defaultHandler(ctx)
	}
}

func helloHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	fmt.Fprintf(ctx, `<doctype html>
		<html>
			<head>
				<title>Test go FastHTTP</title>
			</head>
			<body>
				<h1>SAFECS07</h1>
				<h2>Test go FastHTTP</h2>
			</body>
		</html>`)
}

func defaultHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, World!!!!!!!!!!!!!!!!")
}
