package book

import (
	"github.com/Jezda1337/go-crud/domain/book"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (r *BookRepo) CreateBook(book *book.Book) error {
	if result := r.db.Create(book); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *BookRepo) GetBook(id string) (*book.Book, error) {
	var book book.Book
	var result *gorm.DB
	if result = r.db.First(&book, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

func (r *BookRepo) UpdateBook(id string, data book.Book) (*book.Book, error) {
	var book book.Book
	var result *gorm.DB
	if result = r.db.First(&book, "id = ?", id); result.Error != nil {
		return nil, result.Error
	}

	r.db.Model(&book).Updates(data)

	return &book, nil
}

func (r *BookRepo) GetBooks() ([]*book.Book, error) {
	var books []*book.Book
	if result := r.db.Find(&books); result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}
