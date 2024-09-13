// api/v1/books/handler.go
package books

import (
	"encoding/json"
	"net/http"

	"github.com/fierzahaikkal/neocourse-be-golang/internal/entity"
	borrowModel "github.com/fierzahaikkal/neocourse-be-golang/internal/model/borrow"
	"github.com/fierzahaikkal/neocourse-be-golang/internal/usecase"
	"github.com/gorilla/mux"
)

type BookHandler struct {
	BookUseCase *usecase.BookUseCase
}

func NewBookHandler(bookUseCase *usecase.BookUseCase) *BookHandler {
	return &BookHandler{
		BookUseCase: bookUseCase,
	}
}

func (h *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.Test(w, r) // Handle the GET request
	case http.MethodPost:
		h.GetAllBooksHandler(w, r) // Handle the POST request
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// StoreBookHandler handles storing a new book
func (h *BookHandler) StoreBookHandler(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.BookUseCase.StoreBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// BorrowBookHandler handles borrowing a book
func (h *BookHandler) BorrowBookHandler(w http.ResponseWriter, r *http.Request) {
	var borrowRequest borrowModel.BorrowRequest
	if err := json.NewDecoder(r.Body).Decode(&borrowRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.BookUseCase.BorrowBook(&borrowRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book borrowed successfully"))
}

// GetAllBooksHandler handles fetching all books
func (h *BookHandler) GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, err := h.BookUseCase.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

// GetBookByIDHandler handles fetching a specific book by ID
func (h *BookHandler) GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// if err != nil {
	// 	http.Error(w, "Invalid book ID", http.StatusBadRequest)
	// 	return
	// }

	book, err := h.BookUseCase.FindBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// UpdateBookHandler handles updating a specific book
func (h *BookHandler) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists || id == "" {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book entity.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := h.BookUseCase.UpdateBook(id, &book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}


// DeleteBookHandler handles deleting a specific book
func (h *BookHandler) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] // Use id directly as a string

	err := h.BookUseCase.DeleteBook(id) // Make sure DeleteBook accepts a string
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book deleted successfully"))
}

func (h *BookHandler) Test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
