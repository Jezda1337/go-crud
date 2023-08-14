package book

import (
	"github.com/Jezda1337/go-crud/domain/book"
	"github.com/google/uuid"
)

type BookService struct {
	bookRepo book.BookRepo
}

func NewBookService(bookRepo book.BookRepo) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

func (s *BookService) CreateBook(book *book.Book) error {
	book.ID = uuid.New().String()
	if err := s.bookRepo.CreateBook(book); err != nil {
		return err
	}
	return nil
}

func (s *BookService) GetBook(id string) (*book.Book, error) {
	if book, err := s.bookRepo.GetBook(id); err != nil {
		return nil, err
	} else {
		return book, nil
	}
}

func (s *BookService) UpdateBook(id string, data book.Book) (*book.Book, error) {
	if book, err := s.bookRepo.UpdateBook(id, data); err != nil {
		return nil, err
	} else {
		return book, nil
	}
}

func (s *BookService) GetBooks() ([]*book.Book, error) {
	if books, err := s.bookRepo.GetBooks(); err != nil {
		return nil, err
	} else {
		return books, nil
	}
}
