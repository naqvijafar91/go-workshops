package main

import (
	"encoding/json"
	"net/http"
)

//Note: Only Capitalised fields in a struct are reflected in JSON response
type Response struct {
	Message string    `json:"mess"`
	To      [3]string `json:"recipients"`
	Code    int       `json:"secret_code"`
}

// HTTP handler
func handler(w http.ResponseWriter, r *http.Request) {

	response := Response{
		"Hello World",
		[...]string{
			"Mom",
			"Dad",
			"Grandpa",
		},
		120,
	}

	w.Header().Add("Content-Type", "application/json")
	// Responding with JSON Array
	json.NewEncoder(w).Encode([]Response{response, response})

}

func main() {
	http.HandleFunc("/", handler)
	panic(http.ListenAndServe(":8080", nil))
}
