package borrow

type BorrowResponse struct{
	ID string `json:"id"`
	BorrowedBy string `json:"borrowed_by"`
	BorrowedAt string `json:"borrowed_at"`
}

type BorrowRequest struct{
	ID string `json:"id"`
	BorrowedBy string `json:"borrowed_by"`
	BorrowedAt string `json:"borrowed_at"`
}