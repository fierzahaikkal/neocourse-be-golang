package book

type BookResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Available bool   `json:"available_copies"`
}

type BookRequest struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Available bool   `json:"available_copies"`
}

type BookStore struct {
	Author      string `json:"author" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Year        int    `json:"year" validate:"required"`
	StoredBy    string `json:"storedby" validate:"required"`
	IsBorrowed  bool   `json:"isborrowed" validate:"required"`
	Genre       string `json:"genre" validate:"required"`
	BorrowedBy  string `json:"borrowedby" validate:"required"`
}
