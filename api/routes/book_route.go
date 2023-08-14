package routes

import (
	"github.com/Jezda1337/go-crud/api/handlers"
	"github.com/gin-gonic/gin"
)

func MapBookRoutes(g *gin.RouterGroup, h *handlers.BookHandler) {
	g.POST("", h.CreateBook)
	g.GET(":id", h.GetBook)
	g.PUT(":id", h.UpdateBook)
	g.GET("", h.GetBooks)
}
