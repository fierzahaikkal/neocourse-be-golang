package book

type BookResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Genre		string `json:"genre"`
	Available 	bool   `json:"available"`
	ImageURI	string `json:"image_uri"`
}

type BookStoreRequest struct{
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Genre		string `json:"genre"`
	Available 	bool   `json:"available"`
	ImageURI	string `json:"image_uri"`
	Year        int    `json:"not null"`
	StoredBy    string `json:"stored_by"`
	BorrowedBy  string `json:"borrowed_by"`
}

type BookReturnRequest struct{
	ID			string `json:"id"`
}