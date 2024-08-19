package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
	"strconv"
)

type BookHandlers struct {
	service *service.BookService
}

func NewBookHandlers(service *service.BookService) *BookHandlers {
	return &BookHandlers{service: service}
}


func (h *BookHandlers) GetBooks(w http.ResponseWriter, r *http.Request){
	books, err := h.service.GetBooks()
	if err != nil {
		http.Error(w, "failed to get books", http.StatusInternalServerError)
		return
	}

	if len(books) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandlers) CreateBook(w http.ResponseWriter, r *http.Request){
	var book service.Book

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "failed to decode request", http.StatusBadRequest)
		return
	}

	err = h.service.CreateBook(&book)

	if err != nil {
		http.Error(w, "failed to create book", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) GetBookById(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookById(id)
	if err != nil {
		http.Error(w, "failed to get book", http.StatusBadRequest)
		return
	}

	if book == nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) UpdateBook(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	var book service.Book
	err = json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "failed to decode request", http.StatusBadRequest)
		return
	}

	book.ID = id
	err = h.service.UpdateBook(&book)
	if err != nil {
		http.Error(w, "failed to update book", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandlers) DeleteBook(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid book id", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteBook(id)
	if err != nil {
		println(err)
		http.Error(w, "failed to delete book", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
