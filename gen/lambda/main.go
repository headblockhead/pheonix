package main

import (
	"io"
	"net/http"

	"github.com/a-h/awsapigatewayv2handler"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}))
	awsapigatewayv2handler.ListenAndServe(http.DefaultServeMux)
}
