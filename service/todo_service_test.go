package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestTodoRepoImpl_Create(t *testing.T) {
	repo := &TodoRepoImpl{}
	title := "Belajar Go"
	userID := "012930"

	t.Run("Should create new todo", func(t *testing.T) {
		todo, err := repo.Create(title, userID)

		assert.NoError(t, err)
		assert.NotNil(t, todo)

		assert.NotEmpty(t, todo.ID, "Todo ID should not empty")
		assert.Equal(t, title, todo.Title, "Todo title should match input")
		assert.Equal(t, userID, todo.UserID, "Todo userID should match input")
		assert.False(t, todo.Complete, "Todo complete should be false as a default")

		assert.WithinDuration(t, time.Now(), todo.CreateAt, 5 * time.Second)
		assert.WithinDuration(t, time.Now(), todo.UpdateAt, 5 * time.Second)
	})

	t.Run("Title should not be empty", func(t *testing.T) {
		todo, err := repo.Create(title, userID)

		assert.NotEmpty(t, todo.Title)
		assert.NoError(t, err)
	})
}