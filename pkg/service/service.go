package service

import (
	"api-architecture/models"
	"api-architecture/pkg/repository"
)

type (
	Todo interface {
		Create(note string) error
		Read(id int) (*models.TodoReadOutput, error)
		Update(id int, note string) error
		Delete(id int) error
	}

	Service struct {
		Todo
	}
)

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repo.Todo),
	}
}
