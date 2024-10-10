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

type BookReturnRequest struct{
	ID			string `json:"id"`
}