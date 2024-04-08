package service

import (
	"api-architecture/models"
	"api-architecture/pkg/repository"
	"api-architecture/pkg/transactor"
	"context"
)

type (
	Todo interface {
		Create(note string) error
		Read(id int) (*models.TodoReadOutput, error)
		Update(id int, note string) error
		Delete(id int) error
	}

	Transactor interface {
		WithinTransaction(ctx context.Context, tFunc transactor.Func) (interface{}, error)
	}

	Service struct {
		Todo
		Transactor
	}
)

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Todo:       NewTodoService(repo.Todo),
		Transactor: NewTransactorService(repo.Transactor),
	}
}
