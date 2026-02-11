package service

import (
	"CRUD_GO/internal/model"
	"CRUD_GO/internal/store"
	"errors"
)

type Logger interface {
	log()
}
type Service struct {
	store  store.Store
	logger Logger
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
		//looger: nil,
	}
}

func (s *Service) GetAllBooks() ([]*model.Book, error) {
	return s.store.GetAll()
}
func (s *Service) GetBookByID(id int) (*model.Book, error) {
	return s.store.GetByID(id)
}
func (s *Service) CreateBook(book model.Book) (*model.Book, error) {
	if book.Title == " " {
		return nil, errors.New("Escriba el titulo")
	}
	return s.store.Create(&book)
}

func (s *Service) UpdateBook(id int, book model.Book) (*model.Book, error) {
	if book.Title == " " {
		return nil, errors.New("Escriba el titulo del libro")
	}
	return s.store.Update(id, &book)
}
func (s *Service) Delete(id int) error {
	return s.store.Delete(id)
}
