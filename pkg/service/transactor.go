package service

import (
	"api-architecture/pkg/repository"
	"api-architecture/pkg/transactor"
	"context"
)

type TransactorService struct {
	repo repository.Transactor
}

func NewTransactorService(repo repository.Transactor) *TransactorService {
	return &TransactorService{repo: repo}
}

func (t *TransactorService) WithinTransaction(ctx context.Context, tFunc transactor.Func) (interface{}, error) {
	return t.repo.WithinTransaction(ctx, tFunc)
}
