package main

import (
	"net/http"

	"go-workshops/012-testing/03-testing-http-handler/handler"
)

func main() {
	http.HandleFunc("/", handler.Simple)
	http.ListenAndServe(":8080", nil)
}
