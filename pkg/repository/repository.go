package repository

import (
	"api-architecture/models"
	"github.com/jmoiron/sqlx"
)

type (
	Todo interface {
		Create(note string) error
		Read(id int) (*models.TodoReadOutput, error)
		Update(id int, note string) error
		Delete(id int) error
	}

	Repository struct {
		Todo
	}
)

func NewRepository(connection *sqlx.DB) *Repository {
	return &Repository{
		Todo: NewTodoPostgres(connection),
	}
}
