package repository

import "todo-api/model"

type TodoRepository interface {
	Create(todo *model.Todo) (*model.Todo, error)
	GetAllTodo(userID string) ([]model.Todo, error)
	GetTodoById(userID, todoID string) (*model.Todo, error)
	UpdateTodo(todo *model.Todo) error
}
