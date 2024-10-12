package book

import "github.com/fierzahaikkal/neocourse-be-golang/internal/entity"

func BookMapper(book *entity.Book) *BookResponse {
	return &BookResponse{
		ID:        		book.ID,
		Title:     		book.Title,
		Author:    		book.Author,
		Description: 	book.Description,
		Available: 		book.Available,
		Genre: 			book.Genre,
		ImageURI:		book.ImageURI,
	}
}