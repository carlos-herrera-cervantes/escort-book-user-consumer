package repositories

import (
	"context"
	"escort-book-user-consumer/db"
	"escort-book-user-consumer/models"
	"time"
)

type IDictumRepository interface {
	Create(ctx context.Context, dictum models.Dictum) error
}

type DictumRepository struct {
	Data *db.Data
}

func (r *DictumRepository) Create(ctx context.Context, dictum models.Dictum) error {
	query := "INSERT INTO dictum VALUES ($1, $2, $3, $4, $5, $6, $7);"
	dictum.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		dictum.Id,
		dictum.UserId,
		dictum.FromUser,
		dictum.StatusCategoryId,
		dictum.Comment,
		time.Now().UTC(),
		time.Now().UTC(),
	)

	if err != nil {
		return err
	}

	return nil
}
