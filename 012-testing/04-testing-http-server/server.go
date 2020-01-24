package main

import (
	"net/http"

	"go-workshops/012-testing/04-testing-http-server/handler"
)

func main() {
	http.HandleFunc("/", handler.Simple)
	http.ListenAndServe(":8080", nil)
}
