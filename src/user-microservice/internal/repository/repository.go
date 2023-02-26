package repository

import (
	"context"
	"user-microservice/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User interface {
	FindUserById(ctx context.Context, id string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindAll(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (id uint, err error)
	DeleteUser(ctx context.Context, id string) error
}

type Repository struct {
	User
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
