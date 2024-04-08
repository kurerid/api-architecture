package repository

import (
	"api-architecture/models"
	"api-architecture/pkg/transactor"
	"context"
	"github.com/jmoiron/sqlx"
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

	Repository struct {
		Todo
		Transactor
	}
)

func NewRepository(connection *sqlx.DB) *Repository {
	return &Repository{
		Todo:       NewTodoPostgres(connection),
		Transactor: NewTransactorPostgres(connection),
	}
}
