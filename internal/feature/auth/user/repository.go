package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresRepository struct {
	postgres *sqlx.DB
}

func NewUserRepository(postgresConnection *sqlx.DB) UserRepository {
	return &postgresRepository{
		postgres: postgresConnection,
	}
}

func (r *postgresRepository) GetUserByID(ctx context.Context, id int) (username string, err error) {
	if err := r.postgres.Get(&username, "SELECT username FROM users WHERE id=$1", id); err != nil {
		return username, err
	}

	return username, nil
}
