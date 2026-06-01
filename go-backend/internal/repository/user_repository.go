package repository

import (
	"context"
	"errors"
	"instantanea/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Db *pgxpool.Pool
}

func (r *UserRepository) Find(ctx context.Context, userId int) (model.User, error) {
	var u model.User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email, profile_picture_url, is_active, user_role FROM users WHERE user_id=$1",
		userId,
	).Scan(&u.UserId, &u.Username, &u.Password, &u.Email, &u.ProfilePictureUrl, &u.IsActive, &u.UserRole)

	return u, err
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (model.User, error) {
	var u model.User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email, profile_picture_url, is_active, user_role FROM users WHERE username=$1",
		username,
	).Scan(&u.UserId, &u.Username, &u.Password, &u.Email, &u.ProfilePictureUrl, &u.IsActive, &u.UserRole)

	return u, err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var u model.User

	err := r.Db.QueryRow(
		ctx,
		"SELECT user_id, username, password, email, profile_picture_url, is_active, user_role FROM users WHERE email=$1",
		email,
	).Scan(&u.UserId, &u.Username, &u.Password, &u.Email, &u.ProfilePictureUrl, &u.IsActive, &u.UserRole)

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

func (r *UserRepository) DeleteUser(ctx context.Context, userId int) (error) {
	tag, err := r.Db.Exec(
		ctx,
		"UPDATE users SET is_active=FALSE WHERE user_id=$1",
		userId,
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
		"UPDATE users SET profile_picture_url=$2 WHERE user_id=$1",
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


func (r *UserRepository) GetFollow(ctx context.Context, userId int, followingId int) (model.Follow, error) {

	var follow model.Follow

	err := r.Db.QueryRow(
		ctx,
		"SELECT follower_id, following_id FROM follows WHERE follower_id=$1 AND following_id=$2",
		userId, followingId,
	).Scan(&follow.FollowId, &follow.FollowingId)

	return follow, err
}

func (r *UserRepository) InsertFollow(ctx context.Context, userId int, followingId int) error {
	_, err := r.Db.Exec(
		ctx,
		"INSERT INTO follows (follower_id, following_id) VALUES ($1, $2)",
		userId, followingId,
	)

	return err
}

func (r *UserRepository) RemoveFollow(ctx context.Context, userId int, followingId int) error {
	tag, err := r.Db.Exec(
		ctx,
		"DELETE FROM follows WHERE follower_id=$1 AND following_id=$2",
		userId, followingId,
	)

	if tag.RowsAffected() == 0 {
		return errors.New("0 columnas afectadas")
	}
	
	return err
}

func (r *UserRepository) GetFollowingUsers(ctx context.Context, userId int) ([]model.UserFollowingResponse, error) {

	rows, err := r.Db.Query(ctx, `
    SELECT 
        u.username, u.profile_picture_url, u.user_id, u.user_role,
        TRUE AS following
    FROM follows f
    INNER JOIN users u
        ON u.user_id = f.following_id
    WHERE f.follower_id = $1 AND u.is_active = TRUE
    ORDER BY u.user_id`, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.UserFollowingResponse])

	return users, err	
}