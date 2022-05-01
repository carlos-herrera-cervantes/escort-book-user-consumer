package repositories

import (
	"context"
	"escort-book-user-consumer/db"
	"escort-book-user-consumer/models"
	"fmt"
)

type IStatusCategoryRepository interface {
	GetByField(ctx context.Context, field, value string) (models.StatusCategory, error)
}

type StatusCategoryRepository struct {
	Data *db.Data
}

func (r *StatusCategoryRepository) GetByField(
	ctx context.Context, field, value string,
) (models.StatusCategory, error) {
	query := fmt.Sprintf("SELECT id, name FROM status_category WHERE %s = '%s'", field, value)
	row := r.Data.DB.QueryRowContext(ctx, query)

	var category models.StatusCategory
	err := row.Scan(&category.Id, &category.Name)

	if err != nil {
		return models.StatusCategory{}, err
	}

	return category, nil
}
