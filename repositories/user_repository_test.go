package repositories

import (
	"context"
	"testing"

	"escort-book-user-consumer/db"
	"escort-book-user-consumer/models"

	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryGetByField(t *testing.T) {
	userRepository := UserRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := userRepository.GetByField(context.Background(), "email", "bad@example.com")
		assert.Error(t, err)
	})
}

func TestUserRepositoryCreate(t *testing.T) {
	userRepository := UserRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		ctxWithCancel, cancel := context.WithCancel(context.Background())
		cancel()
		err := userRepository.Create(ctxWithCancel, models.User{})
		assert.Error(t, err)
	})
}
