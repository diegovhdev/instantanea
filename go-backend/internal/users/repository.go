package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Db *pgxpool.Pool
}

func (r *Repository) Find(user_id int, ctx context.Context) (User, error) {
	var u User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email FROM users WHERE user_id=$1",
		user_id,
	).Scan(&u.Id, &u.Username, &u.Password, &u.Email)

	return u, err
}

func (r *Repository) FindByUsername(username string, ctx context.Context) (User, error) {
	var u User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email FROM users WHERE username=$1",
		username,
	).Scan(&u.Id, &u.Username, &u.Password, &u.Email)

	return u, err
}

func (r *Repository) FindByEmail(email string, ctx context.Context) (User, error) {
	var u User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email FROM users WHERE email=$1",
		email,
	).Scan(&u.Id, &u.Username, &u.Password, &u.Email)

	return u, err
}

func (r *Repository) Insert(user User, ctx context.Context) (error) {
	_, err := r.Db.Exec(
		ctx,
		"INSERT INTO users (username, password, email) VALUES ($1, $2, $3)",
		user.Username, 
		user.Password, 
		user.Email,
	)

	return err
}



func (r *Repository) Delete(user_id int, ctx context.Context) (error) {
		result, err := r.Db.Exec(
		ctx,
		"DELETE FROM users WHERE user_id == $1",
		user_id,
	)

	if result.RowsAffected() == 0 {
		return fmt.Errorf("USER DOESN'T EXIST")
	}

	return err
}