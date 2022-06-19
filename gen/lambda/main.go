package main

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/a-h/awsapigatewayv2handler"
	"github.com/headblockhead/phoenix/handler"
)

func main() {
	awsapigatewayv2handler.ListenAndServe(gziphandler.GzipHandler(http.HandlerFunc(handler.Handle)))
}
