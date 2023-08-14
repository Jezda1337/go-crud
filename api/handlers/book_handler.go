package handlers

import (
	"net/http"

	"github.com/Jezda1337/go-crud/domain/book"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService book.BookService
}

func NewBookHandler(bookService book.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (h *BookHandler) CreateBook(ctx *gin.Context) {
	var book book.Book
	if err := ctx.BindJSON(&book); err != nil {
		panic(err)
	}

	if err := h.bookService.CreateBook(&book); err != nil {
		panic(err)
	}
}

func (h *BookHandler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	if book, err := h.bookService.GetBook(id); err != nil {
		panic(err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"book": *book,
		})
	}
}

func (h *BookHandler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var input book.Book

	if err := ctx.ShouldBindJSON(&input); err != nil {
		panic(err)
	}

	if book, err := h.bookService.UpdateBook(id, input); err != nil {
		panic(err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"book": book,
		})
	}
}

func (h *BookHandler) GetBooks(ctx *gin.Context) {
	if books, err := h.bookService.GetBooks(); err != nil {
		panic(err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"books": books,
		})
	}
}
