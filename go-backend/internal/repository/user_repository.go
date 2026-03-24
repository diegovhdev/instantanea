package repository

import (
	"context"
	"instantanea/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Db *pgxpool.Pool
}

func (r *UserRepository) Find(ctx context.Context, user_id int) (model.User, error) {
	var u model.User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email FROM users WHERE user_id=$1",
		user_id,
	).Scan(&u.UserId, &u.Username, &u.Password, &u.Email)

	return u, err
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var u model.User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email FROM users WHERE username=$1",
		username,
	).Scan(&u.UserId, &u.Username, &u.Password, &u.Email)

	return u, err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var u model.User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email FROM users WHERE email=$1",
		email,
	).Scan(&u.UserId, &u.Username, &u.Password, &u.Email)

	return u, err
}

func (r *UserRepository) Insert(ctx context.Context, user model.User) (error) {
	_, err := r.Db.Exec(
		ctx,
		"INSERT INTO users (username, password, email) VALUES ($1, $2, $3)",
		user.Username, 
		user.Password, 
		user.Email,
	)

	return err
}
