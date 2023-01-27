package repositories

import (
	"context"
	"testing"

	"escort-book-user-consumer/db"

	"github.com/stretchr/testify/assert"
)

func TestStatusCategoryRepositoryGetByField(t *testing.T) {
	statusCategoryRepository := StatusCategoryRepository{
		Data: db.NewPostgresClient(),
	}

	t.Run("Should return error", func(t *testing.T) {
		_, err := statusCategoryRepository.GetByField(context.Background(), "name", "bad")
		assert.Error(t, err)
	})
}
