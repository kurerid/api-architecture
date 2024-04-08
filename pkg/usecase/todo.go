package usecase

import "api-architecture/models"

func (u *Usecase) CreateTodo(input models.TodoCreateInput) error {
	return u.services.Todo.Create(input.Note)
}

func (u *Usecase) ReadTodo(todoID int) (*models.TodoReadOutput, error) {
	return u.services.Todo.Read(todoID)
}

func (u *Usecase) UpdateTodo(input models.TodoUpdateInput) error {
	return u.services.Todo.Update(input.Id, input.Note)
}

func (u *Usecase) DeleteTodo(todoID int) error {
	return u.services.Todo.Delete(todoID)
}
