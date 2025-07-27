package service

import (
	"errors"
	"time"
	"todo-api/model"
	"todo-api/repository"

	"github.com/google/uuid"
)

type TodoRepoImpl struct {
	Repo *repository.TodoRepository
}

func (t *TodoRepoImpl) Create(title, userID string) (*model.Todo, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	todo := &model.Todo{
		ID: uuid.NewString(),
		Title: title,
		Complete: false,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		UserID: userID,
	}

	return todo, nil
}