package repository

import (
	"context"

	"user-microservice/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx, "SELECT id, email, password FROM public.user WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Password)
	return &user, err
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	rows, err := r.db.Query(ctx, "SELECT id, email, password FROM public.user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) AddUser(user *model.User) error {
	_, err := r.db.Exec(context.Background(), "INSERT INTO public.user(email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow(ctx, "SELECT id, email, password FROM public.user WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password)
	return &user, err
}
