package transport

import (
	"CRUD_GO/internal/model"
	"CRUD_GO/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	service *service.Service
}

func New(s *service.Service) *BookHandler {
	return &BookHandler{
		service: s,
	}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		books, err := h.service.GetAllBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)

	case http.MethodPost:
		var book model.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		created, err := h.service.CreateBook(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}

func (h *BookHandler) HandleBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "no lo encontre", http.StatusBadRequest)
	}

	switch r.Method {
	case http.MethodGet:
		book, err := h.service.GetBookByID(id)
		if err != nil {
			http.Error(w, "No lo encontramos", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)

	case http.MethodPut:
		var book model.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, "entrada invalida", http.StatusBadRequest)
			return
		}
		update, err := h.service.UpdateBook(id, book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(update)

	case http.MethodDelete:
		if err := h.service.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Metodo no disponible", http.StatusMethodNotAllowed)
	}
}
