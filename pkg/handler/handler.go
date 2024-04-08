package handler

import (
	"api-architecture/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecases *usecase.Usecase
}

func NewHandler(usecases *usecase.Usecase) *Handler {
	return &Handler{usecases: usecases}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	todo := router.Group("/todo")
	{
		todo.POST("", h.createTodo)
		todo.GET("/:todoID", h.readTodo)
		todo.PUT("", h.updateTodo)
		todo.DELETE("/:todoID", h.deleteTodo)
	}

	return router
}
