package main

import (
	"testing"
)

func TestBasic(t *testing.T) {

	books := []Book{Book{
		Name:     "Book 1",
		AuthorID: 1,
		ID:       1,
	}, Book{
		Name:     "Book 2",
		AuthorID: 1,
		ID:       2,
	}, Book{
		Name:     "Book 3",
		AuthorID: 2,
		ID:       3,
	}}

	authors := []Author{Author{
		Name: "Author 1",
		ID:   1,
	}, Author{
		Name: "Author 2",
		ID:   2,
	}}

	service := NewCombinationService()
	response := service.GenerateResponse(books, authors)
	if len(response) != 3 {
		t.Errorf("Incorrect length - Expected %d, found %d", 3, len(response))
	}
	if response[0].AuthorID != 0 {
		t.Error("Author ID should be zero so it could be omitted from response")
	}
}
