package book

type BookStoreRequest struct{
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Genre		string `json:"genre"`
	ImageURI	string `json:"image_uri"`
	Year        int    `json:"year"`
}

type UpdateBookRequest struct {
	Author      *string `json:"author"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Year        *int    `json:"year"`
	Genre       *string `json:"genre"`
	ImageURI    *string `json:"image_uri"`
}

type GetBookResponse struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Genre       string `json:"genre"`
	StoredBy 	string `json:"stored_by"`
	ImageURI    string `json:"image_uri"`
}