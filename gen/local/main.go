package main

import (
	"net/http"

	"github.com/headblockhead/phoenix/handler"
)

func main() {
	// http.ListenAndServe("0.0.0.0:8000", gziphandler.GzipHandler(http.HandlerFunc(handler.Handle)))
	http.ListenAndServe("0.0.0.0:8000", http.HandlerFunc(handler.Handle))
}
