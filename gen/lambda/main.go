package main

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/headblockhead/phoenix/handler"
)

func main() {
	h := gziphandler.GzipHandler(http.HandlerFunc(handler.Handle))
	adapter := httpadapter.New(h)
	lambda.Start(adapter.ProxyWithContext)
}
