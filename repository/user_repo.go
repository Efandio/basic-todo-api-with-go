package repository

import "todo-api/model"

type UserRepository interface {
	Username() error
	Email() error
	Password() error
	Todo() []model.Todo
}
