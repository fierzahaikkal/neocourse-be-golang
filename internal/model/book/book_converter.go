package book

import "github.com/fierzahaikkal/neocourse-be-golang/internal/entity"

func BookMapper(book *entity.Book) *BookResponse {
	return &BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Available: book.IsBorrowed,
	}
}
