package repository

import (
	"context"
	"instantanea/internal/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	Db *pgxpool.Pool
}

func (r *PostRepository) Insert(ctx context.Context, post model.Post) (error) {
	_, err := r.Db.Exec(
		ctx,
		"INSERT INTO posts (user_id, url, public_id, text) VALUES ($1, $2, $3, $4)",
		post.UserId, 
		post.Url, 
		post.PublicId,
		post.Text,
	)

	return err
}

func (r *PostRepository) GetMany(ctx context.Context, limit int, offset int) ([]model.PostResponse, error) {

	rows, err := r.Db.Query(ctx, `
	SELECT p.post_id, p.user_id, u.username, p.text, p.url 
	FROM posts as p
	JOIN users AS u ON p.user_id = u.user_id
	ORDER BY p.post_id DESC
	LIMIT $1 OFFSET $2;`, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.PostResponse])

	return posts, err
}


func (r *PostRepository) GetManyOrderedById(ctx context.Context, stopId int) ([]model.PostResponse, error) {
	rows, err := r.Db.Query(ctx, `
	SELECT p.post_id, p.user_id, u.username, u.profile_picture_url, p.text, p.url, p.votes
	FROM posts as p
	JOIN users AS u ON p.user_id = u.user_id
	WHERE u.is_active = TRUE AND p.post_id >$1
	ORDER BY p.post_id DESC;`, stopId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.PostResponse])

	return posts, err
}

func (r *PostRepository) GetManyOrderedByVotes(ctx context.Context, stopId int) ([]model.PostResponse, error) {
	rows, err := r.Db.Query(ctx, `
	SELECT p.post_id, p.user_id, u.username, u.profile_picture_url, p.text, p.url, p.votes
	FROM posts as p
	JOIN users AS u ON p.user_id = u.user_id
	WHERE u.is_active = TRUE AND p.post_id >$1
	ORDER BY p.votes, p.post_id DESC;`, stopId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.PostResponse])

	return posts, err
}

func (r *PostRepository) InsertVote(postId int, userId int) {
}

func (r *PostRepository) GetManyOrderedByFavorites(ctx context.Context, stopId int, id int) ([]model.PostResponse, error) {
	rows, err := r.Db.Query(ctx, `
	SELECT p.post_id, p.user_id, u.username, u.profile_picture_url, p.text, p.url, p.votes
	FROM posts as p
	JOIN users AS u ON p.user_id = u.user_id
	JOIN votes AS v ON p.post_id = v.post_id
	WHERE u.is_active = TRUE AND p.post_id > $1 AND v.user_id = $2
	ORDER BY p.votes DESC;`, stopId, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.PostResponse])

	return posts, err
}