package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {

}

func testHandler(ctx *fasthttp.RequestCtx) {
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

}
