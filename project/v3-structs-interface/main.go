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

type BookRepository interface {
	GetAll() ([]Book, error)
	Create(book *Book) error
}

type AuthorRepository interface {
	GetAll() ([]Author, error)
	Create(author *Author) error
}

type Handler struct {
	bookRepository   BookRepository
	authorRepository AuthorRepository
}

func (h *Handler) SaveBook(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var book *Book
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
	// Save
	err = h.bookRepository.Create(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	// Covert books map to slice
	response, err := h.bookRepository.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	// Responding with JSON Array
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) SaveAuthor(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var author *Author
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
	err = h.authorRepository.Create(author)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(author)
}

func (h *Handler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	// Covert books map to slice
	response, err := h.authorRepository.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Add("Content-Type", "application/json")
	// Responding with JSON Array
	json.NewEncoder(w).Encode(response)
}

// ListenAndServe start listening for client connections on given port
func main() {
	api := &Handler{NewMemoryBackedBookRepository(), NewMemoryBackedAuthorRepository()}
	http.HandleFunc("/authors", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			api.GetAllAuthors(w, r)
		case http.MethodPost:
			api.SaveAuthor(w, r)
		default:
			// Give an error message.
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request method."))
		}
	})
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			api.GetAllBooks(w, r)
		case http.MethodPost:
			api.SaveBook(w, r)
		default:
			// Give an error message.
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request method."))
		}
	})
	http.ListenAndServe(":8080", nil)
}
