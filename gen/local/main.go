package main

import (
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/headblockhead/phoenix/handler"
)

func main() {
	http.ListenAndServe("0.0.0.0:8000", gziphandler.GzipHandler(http.HandlerFunc(handler.Handle)))
}
