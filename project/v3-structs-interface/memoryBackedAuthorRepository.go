package main

import "errors"

type MemoryBackedAuthorRepository struct {
	authors map[string]Author
}

func (repo *MemoryBackedAuthorRepository) GetAll() ([]Author, error) {
	response := make([]Author, 0)
	for _, v := range repo.authors {
		response = append(response, v)
	}
	return response, nil
}

func (repo *MemoryBackedAuthorRepository) Create(author *Author) error {
	// Check for duplicacy
	if _, ok := repo.authors[author.Name]; ok {
		return errors.New("Duplicate Author Found")
	}
	// Generate an ID
	author.ID = len(repo.authors) + 1
	repo.authors[author.Name] = *author
	return nil
}

// Constructor Function
func NewMemoryBackedAuthorRepository() AuthorRepository {
	return &MemoryBackedAuthorRepository{make(map[string]Author)}
}
