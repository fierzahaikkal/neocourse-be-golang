package borrow

type BorrowRequest struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Genre		string `json:"genre"`
	BorrowedBy 	string   `json:"borrowedBy"`
}