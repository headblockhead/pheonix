package main

import (
	"net/http"

	"github.com/headblockhead/phoenix/handler"
)

func main() {
	http.ListenAndServe("localhost:8000", http.HandlerFunc(handler.Handle))
}
