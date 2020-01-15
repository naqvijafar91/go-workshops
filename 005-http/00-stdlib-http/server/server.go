package main

import (
	"net/http"
)

// handler is used to write data into response body
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func help(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Use / endpoint! "))
}

// ListenAndServe start listening for client connections on given port
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Serve the resource.
			help(w, r)
		default:
			// Give an error message.
			w.Write([]byte("Invalid request method. This endpoint only supports GET"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
