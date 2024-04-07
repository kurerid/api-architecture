package service

import (
	"api-architecture/models"
	"api-architecture/pkg/repository"
)

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) Create(note string) error {
	return s.repo.Create(note)
}

func (s *TodoService) Read(id int) (*models.TodoReadOutput, error) {
	return s.repo.Read(id)
}

func (s *TodoService) Update(id int, note string) error {
	return s.repo.Update(id, note)
}

func (s *TodoService) Delete(id int) error {
	return s.repo.Delete(id)
}
