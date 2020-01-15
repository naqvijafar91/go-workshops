package main

type CombinationServiceImpl struct {
}

func (cb *CombinationServiceImpl) GenerateResponse(books []Book, authors []Author) []CombinedResponse {
	combinedResponse := make([]CombinedResponse, 0)
	// Iterate over books
	for _, book := range books {
		// Find author for every book
		for _, author := range authors {
			if author.ID == book.AuthorID {
				book.AuthorID = 0
				singleResponse := CombinedResponse{
					book,
					author,
				}
				combinedResponse = append(combinedResponse, singleResponse)
			}
		}
	}
	return combinedResponse
}

func NewCombinationService() CombinationService {
	return &CombinationServiceImpl{}
}
