package main

import (
	"github.com/Jezda1337/go-crud/api/handlers"
	"github.com/Jezda1337/go-crud/api/routes"
	"github.com/Jezda1337/go-crud/db"
	"github.com/Jezda1337/go-crud/infra/book"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	apiRoute := r.Group("/api/v1")
	bookRoute := apiRoute.Group("/books")
	bookRepo := book.NewBookRepo(db)
	bookService := book.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)

	routes.MapBookRoutes(bookRoute, bookHandler)


	r.Run()
}
