package repositories

import (
	"context"
	"escort-book-user-consumer/db"
	"escort-book-user-consumer/models"
	"fmt"
	"time"
)

type IUserRepository interface {
	GetByField(ctx context.Context, field, value string) (models.User, error)
	Create(ctx context.Context, user models.User) error
}

type UserRepository struct {
	Data *db.Data
}

func (r *UserRepository) GetByField(ctx context.Context, field, value string) (models.User, error) {
	query := fmt.Sprintf("SELECT email FROM user WHERE %s = %s", field, value)
	row := r.Data.DB.QueryRowContext(ctx, query)

	var user models.User
	err := row.Scan(&user.Email)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user models.User) error {
	query := "INSERT INTO user VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"
	user.SetDefaultValues()

	_, err := r.Data.DB.ExecContext(
		ctx,
		query,
		user.Id,
		user.UserId,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Deleted,
		time.Now().UTC(),
		time.Now().UTC(),
	)

	if err != nil {
		return err
	}

	return nil
}
