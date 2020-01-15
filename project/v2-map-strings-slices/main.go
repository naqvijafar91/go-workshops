package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	ID       int    `json:"id"` // Auto
	Name     string `json:"name"`
	AuthorID string `json:"authorId"`
}

type Author struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

var books map[string]Book
var authors map[string]Author

func SaveBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	err := json.Unmarshal(reqBody, &book)
	// Error handling Read 10-errors-panics.md
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse request body"))
		return
	}
	if len(book.Name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book name cannot be empty"))
		return
	}
	// Check for duplicacy
	if _, ok := books[book.Name]; ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Duplicate Book"))
		return
	}
	// Generate an ID
	book.ID = len(books) + 1
	// Save
	books[book.Name] = book
	json.NewEncoder(w).Encode(book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	// Covert books map to slice
	response := make([]Book, 0)
	for _, v := range books {
		response = append(response, v)
	}
	w.Header().Add("Content-Type", "application/json")
	// Responding with JSON Array
	json.NewEncoder(w).Encode(response)
}

func SaveAuthor(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var author Author
	err := json.Unmarshal(reqBody, &author)
	// Error handling Read 10-errors-panics.md
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse request body"))
		return
	}
	if len(author.Name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Author name cannot be empty"))
		return
	}
	// Check for duplicacy
	if _, ok := authors[author.Name]; ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Duplicate Author"))
		return
	}
	// Generate an ID
	author.ID = len(authors) + 1
	// Save
	authors[author.Name] = author
	json.NewEncoder(w).Encode(author)
}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	// Covert books map to slice
	response := make([]Author, 0)
	for _, v := range authors {
		response = append(response, v)
	}
	w.Header().Add("Content-Type", "application/json")
	// Responding with JSON Array
	json.NewEncoder(w).Encode(response)
}

// ListenAndServe start listening for client connections on given port
func main() {
	books = make(map[string]Book)
	authors = make(map[string]Author)

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
