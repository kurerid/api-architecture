package handler

import (
	"api-architecture/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createTodo(c *gin.Context) {
	var input models.TodoCreateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.services.Todo.Create(input.Note); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) readTodo(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("todoID"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	output, err := h.services.Todo.Read(todoID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &output)
}

func (h *Handler) updateTodo(c *gin.Context) {
	var input models.TodoUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.services.Todo.Update(input.Id, input.Note); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteTodo(c *gin.Context) {
	todoID, err := strconv.Atoi(c.Param("todoID"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err = h.services.Todo.Delete(todoID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
