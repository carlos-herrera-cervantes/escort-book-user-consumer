package repositories

import (
	"context"
	"fmt"
	"time"

	"escort-book-user-consumer/db"
	"escort-book-user-consumer/models"
)

//go:generate mockgen -destination=./mocks/iuser_repository.go -package=mocks --build_flags=--mod=mod . IUserRepository
type IUserRepository interface {
	GetByField(ctx context.Context, field, value string) (models.User, error)
	Create(ctx context.Context, user models.User) error
}

type UserRepository struct {
	Data *db.PostgresClient
}

func (r *UserRepository) GetByField(ctx context.Context, field, value string) (models.User, error) {
	query := fmt.Sprintf(`SELECT user_id FROM "user" WHERE %s = '%s'`, field, value)
	row := r.Data.UserDB.QueryRowContext(ctx, query)

	var user models.User

	if err := row.Scan(&user.UserId); err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Create(ctx context.Context, user models.User) error {
	query := `INSERT INTO "user" VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	user.SetDefaultValues()

	if _, err := r.Data.UserDB.ExecContext(
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
	); err != nil {
		return err
	}

	return nil
}
