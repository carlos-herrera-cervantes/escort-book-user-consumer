package repositories

import (
	"context"
	"testing"

	"escort-book-user-consumer/db"
	"escort-book-user-consumer/models"

	"github.com/stretchr/testify/assert"
)

func TestDictumRepositoryCreate(t *testing.T) {
	dictumRepository := DictumRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		err := dictumRepository.Create(context.Background(), models.Dictum{})
		assert.Error(t, err)
	})
}
