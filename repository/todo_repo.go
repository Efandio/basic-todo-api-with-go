package repository

import "todo-api/model"

type TodoRepository interface {
	Create(title, userID string) (*model.Todo, error)
	GetAll(userID string) ([]model.Todo, error)
	GetById(userID, todoID string) (*model.Todo, error)
	Update(todo *model.Todo) error
	Delete(todoID string) error
}
