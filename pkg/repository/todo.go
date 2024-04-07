package repository

import (
	"api-architecture/models"
	"database/sql"
	"errors"
	"github.com/execaus/exloggo"
	"github.com/jmoiron/sqlx"
)

type TodoPostgres struct {
	db *sqlx.DB
}

func NewTodoPostgres(connection *sqlx.DB) *TodoPostgres {
	return &TodoPostgres{db: connection}
}

func (r *TodoPostgres) Create(note string) error {
	_, err := r.db.Exec(`INSERT INTO "Todo"("note") VALUES ($1)`, note)
	if err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *TodoPostgres) Read(id int) (*models.TodoReadOutput, error) {
	var output models.TodoReadOutput
	err := r.db.Get(&output.Todo, `SELECT * FROM "Todo" WHERE id = $1`, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		exloggo.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (r *TodoPostgres) Update(id int, note string) error {
	_, err := r.db.Exec(`UPDATE "Todo" SET note = $1 WHERE id =$2`, note, id)
	if err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}

func (r *TodoPostgres) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM "Todo" WHERE id = $1`, id)
	if err != nil {
		exloggo.Error(err.Error())
		return err
	}

	return nil
}
