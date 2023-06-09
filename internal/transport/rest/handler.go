package rest

import (
	"github.com/go-playground/validator/v10"
	"net/http"

	"bibliotekaProject/internal/service/library"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	validate  *validator.Validate
	libraries *library.Service
}

func NewHandler(libraries *library.Service) *Handler {
	return &Handler{
		validate:  validator.New(),
		libraries: libraries,
	}
}

func (h *Handler) Init() http.Handler {
	// Init a new router instance
	router := gin.Default()

	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Register a new routes
	api := router.Group("/api")

	authors := api.Group("/authors")
	{
		authors.POST("", h.CreateAuthor)
		authors.GET("", h.GetAuthors)
		authors.GET("/:id/books", h.GetBooksByAuthorID)
		authors.GET("/:id", h.GetAuthor)
		authors.PUT("/:id", h.UpdateAuthor)
		authors.DELETE("/:id", h.DeleteAuthor)
	}

	books := api.Group("/books")
	{
		books.POST("", h.CreateBook)
		books.GET("", h.GetBooks)
		books.GET("/:id", h.GetBook)
		books.PUT("/:id", h.UpdateBook)
		books.DELETE("/:id", h.DeleteBook)
	}

	members := api.Group("/members")
	{
		members.POST("", h.CreateMember)
		members.GET("", h.GetMembers)
		members.GET("/:id", h.GetMember)
		members.GET("/:id/books", h.GetMemberIdBooks)
		members.PUT("/:id", h.UpdateMember)
		members.DELETE("/:id", h.DeleteMember)
	}

	return router
}
