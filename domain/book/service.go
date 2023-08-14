package book

type BookService interface {
	CreateBook(book *Book) error
	GetBook(id string) (*Book, error)
	UpdateBook(id string, book Book) (*Book, error)
	GetBooks() ([]*Book, error)
}
