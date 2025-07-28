package service

import (
	"database/sql"
	"errors"
	"time"
	"todo-api/model"
	"todo-api/repository"

	"github.com/google/uuid"
)

type TodoRepoImpl struct {
	DB *sql.DB
}


func NewTodoRepoImpl(db *sql.DB) repository.TodoRepository {
	return &TodoRepoImpl{
		DB: db,
	}
}

// Create implements repository.TodoRepository.
func (t *TodoRepoImpl) Create(title, userID string) (*model.Todo, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	todo := &model.Todo{
		ID:       uuid.NewString(),
		Title:    title,
		Complete: false,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		UserID:   userID,
	}

	return todo, nil
}

// Delete implements repository.TodoRepository.
func (t *TodoRepoImpl) Delete(todoID string) error {
	panic("unimplemented")
}

// GetAll implements repository.TodoRepository.
func (t *TodoRepoImpl) GetAll(userID string) ([]model.Todo, error) {
	panic("unimplemented")
}

// GetById implements repository.TodoRepository.
func (t *TodoRepoImpl) GetById(userID string, todoID string) (*model.Todo, error) {
	panic("unimplemented")
}

// Update implements repository.TodoRepository.
func (t *TodoRepoImpl) Update(todo *model.Todo) error {
	panic("unimplemented")
}
