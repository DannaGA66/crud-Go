package service

import (
	"CRUD_GO/internal/model"
	"CRUD_GO/internal/store"
	"errors"
)

type Service struct {
	store store.Store
}

func New(s store.Store) *Service {
	return &Service{
		store: s,
	}
}

func (s *Service) GetAllBooks() ([]*model.Book, error) {
	books, err := s.store.GetAll()

	if err != nil {
		return nil, err
	}
	return books, nil
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
