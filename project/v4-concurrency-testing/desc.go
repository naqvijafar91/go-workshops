package main

/*
Add a new endpoint:

/books-authors - GET: Fetch all books along with their authors in the below format:
[
  {
    "id": 1,
    "name": "My Book",
    "author":  {
    	"name": "Author 1",
    	"id": 1
  	}
  },
  {
    "id": 2,
    "name": "My Book 2",
    "author":  {
    	"name": "Author 1",
    	"id": 1
  	}
  },
  {
    "id": 3,
    "name": "My Book 3",
    "author":  {
    	"name": "Author 2",
    	"id": 2
  	}
  }
]

Use the following approach:
1. Fetch data from both repositories concurrently using 2 goroutines.
2. Create a servie using a struct which combines them together according to the above response.
3. Write a unit test for the above service.
*/
