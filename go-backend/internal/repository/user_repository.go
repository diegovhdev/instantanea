package repository

import (
	"context"
	"errors"
	"instantanea/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Db *pgxpool.Pool
}

func (r *UserRepository) Find(ctx context.Context, userId int) (model.User, error) {
	var u model.User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email, profile_image_url FROM users WHERE user_id=$1",
		userId,
	).Scan(&u.UserId, &u.Username, &u.Password, &u.Email, &u.ProfileImageUrl)

	return u, err
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var u model.User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email, profile_image_url FROM users WHERE username=$1",
		username,
	).Scan(&u.UserId, &u.Username, &u.Password, &u.Email, &u.ProfileImageUrl)

	return u, err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var u model.User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email, profile_image_url FROM users WHERE email=$1",
		email,
	).Scan(&u.UserId, &u.Username, &u.Password, &u.Email, &u.ProfileImageUrl)

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

func (r *UserRepository) UpdateUsername(ctx context.Context, userId int, username string) (error) {
	tag, err := r.Db.Exec(
		ctx,
		"UPDATE users SET username=$2 WHERE user_id=$1",
		userId,
		username,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("0 columnas afectadas")
	}

	return nil
}

func (r *UserRepository) UpdateEmail(ctx context.Context, userId int, email string) (error) {
	tag, err := r.Db.Exec(
		ctx,
		"UPDATE users SET email=$2 WHERE user_id=$1",
		userId,
		email,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("0 columnas afectadas")
	}

	return nil
}

func (r *UserRepository) UpdateProfilePicture(ctx context.Context, userId int, url string) (error) {
	tag, err := r.Db.Exec(
		ctx,
		"UPDATE users SET profile_image_url=$2 WHERE user_id=$1",
		userId,
		url,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("0 columnas afectadas")
	}

	return nil
}

func (r *UserRepository) UpdatePassword(ctx context.Context, userId int, password string) (error) {
	tag, err := r.Db.Exec(
		ctx,
		"UPDATE users SET password=$2 WHERE user_id=$1",
		userId,
		password,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("0 columnas afectadas")
	}

	return nil
}