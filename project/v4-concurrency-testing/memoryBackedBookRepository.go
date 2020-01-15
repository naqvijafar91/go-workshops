package main

import "errors"

type MemoryBackedBookRepository struct {
	books map[string]Book
}

func (repo *MemoryBackedBookRepository) GetAll() ([]Book, error) {
	response := make([]Book, 0)
	for _, v := range repo.books {
		response = append(response, v)
	}
	return response, nil
}

func (repo *MemoryBackedBookRepository) Create(book *Book) error {
	// Check for duplicacy
	if _, ok := repo.books[book.Name]; ok {
		return errors.New("Duplicate book Found")
	}
	// Generate an ID
	book.ID = len(repo.books) + 1
	repo.books[book.Name] = *book
	return nil
}

// Constructor Function
func NewMemoryBackedBookRepository() BookRepository {
	return &MemoryBackedBookRepository{make(map[string]Book)}
}
