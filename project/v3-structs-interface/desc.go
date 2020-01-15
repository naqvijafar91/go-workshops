package main

/*
Update the app to use interface and its implementation for accessing the in memory structures which we created

1. Define the following interface which will be used by the handlers:

type BookRepository interface {
	GetAll() ([]Book, error)
	Create(book Book) error
}

type AuthorRepository interface {
	GetAll() ([]Author, error)
	Create(author Author) error
}

2. Define a Handler struct and use the existing 4 handler functions as methods of that struct.

3. Use Dependency injection concept to pass the implementation of both the repositories to the Handler, this will be done in the
	main function.
*/