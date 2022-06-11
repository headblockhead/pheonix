package main

import (
	"net/http"

	"github.com/a-h/awsapigatewayv2handler"
	"github.com/headblockhead/phoenix/handler"
)

func main() {
	awsapigatewayv2handler.ListenAndServe(http.HandlerFunc(handler.Handle))
}
