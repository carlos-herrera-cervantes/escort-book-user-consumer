package repositories

import (
	"context"
	"fmt"

	"escort-book-user-consumer/db"
	"escort-book-user-consumer/models"
)

//go:generate mockgen -destination=./mocks/istatus_category_repository.go -package=mocks --build_flags=--mod=mod . IStatusCategoryRepository
type IStatusCategoryRepository interface {
	GetByField(ctx context.Context, field, value string) (models.StatusCategory, error)
}

type StatusCategoryRepository struct {
	Data *db.PostgresClient
}

func (r *StatusCategoryRepository) GetByField(
	ctx context.Context, field, value string,
) (models.StatusCategory, error) {
	query := fmt.Sprintf("SELECT id, name FROM status_category WHERE %s = '%s'", field, value)
	row := r.Data.UserDB.QueryRowContext(ctx, query)

	var category models.StatusCategory

	if err := row.Scan(&category.Id, &category.Name); err != nil {
		return category, err
	}

	return category, nil
}
