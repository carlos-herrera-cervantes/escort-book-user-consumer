package repositories

import (
	"context"
	"time"

	"escort-book-user-consumer/db"
	"escort-book-user-consumer/models"
)

//go:generate mockgen -destination=./mocks/idictum_repository.go -package=mocks --build_flags=--mod=mod . IDictumRepository
type IDictumRepository interface {
	Create(ctx context.Context, dictum models.Dictum) error
}

type DictumRepository struct {
	Data *db.PostgresClient
}

func (r *DictumRepository) Create(ctx context.Context, dictum models.Dictum) error {
	query := "INSERT INTO dictum VALUES ($1, $2, $3, $4, $5, $6, $7);"
	dictum.SetDefaultValues()

	if _, err := r.Data.UserDB.ExecContext(
		ctx,
		query,
		dictum.Id,
		dictum.UserId,
		dictum.FromUser,
		dictum.StatusCategoryId,
		dictum.Comment,
		time.Now().UTC(),
		time.Now().UTC(),
	); err != nil {
		return err
	}

	return nil
}
