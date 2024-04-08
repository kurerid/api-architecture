package repository

import (
	"api-architecture/pkg/transactor"
	"context"
	"github.com/jmoiron/sqlx"
)

type TransactorPostgres struct {
	db *sqlx.DB
}

func NewTransactorPostgres(db *sqlx.DB) *TransactorPostgres {
	return &TransactorPostgres{db: db}
}

func (r *TransactorPostgres) WithinTransaction(ctx context.Context, tFunc transactor.Func) (interface{}, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	output, err := tFunc(injectTx(ctx, tx))
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()
	return output, nil
}

func injectTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, transactor.Key{}, tx)
}

func extractTx(ctx context.Context) *sqlx.Tx {
	if tx, ok := ctx.Value(transactor.Key{}).(*sqlx.Tx); ok {
		return tx
	}
	return nil
}
