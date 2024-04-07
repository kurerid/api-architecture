package handler

import (
	"api-architecture/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
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
