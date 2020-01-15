package main

import (
	"net/http"
)

func SaveBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Book Saved"))
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Returning all books"))
}

func SaveAuthor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Author Saved"))
}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Returning all authors"))
}

// ListenAndServe start listening for client connections on given port
func main() {
	http.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllAuthors(w, r)
		case http.MethodPost:
			SaveAuthor(w, r)
		default:
			// Give an error message.
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request method."))
		}
	})
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetAllBooks(w, r)
		case http.MethodPost:
			SaveBook(w, r)
		default:
			// Give an error message.
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request method."))
		}
	})
	http.ListenAndServe(":8080", nil)
}
