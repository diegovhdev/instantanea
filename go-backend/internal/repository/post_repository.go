package repository

import (
	"context"
	"errors"
	"instantanea/internal/model"
	"log"

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


func (r *PostRepository) GetManyOrderedById(ctx context.Context, userId int, stopId int) ([]model.PostResponse, error) {
	rows, err := r.Db.Query(ctx, `
	SELECT
    	p.post_id, p.user_id, p.url, p.text, p.votes, 
    	u.username,
    	u.profile_picture_url,
    	CASE WHEN v.user_id IS NOT NULL THEN TRUE ELSE FALSE END AS voted,
		CASE WHEN f.follower_id IS NOT NULL THEN TRUE ELSE FALSE END AS following
	FROM posts p
	LEFT JOIN votes v
   		ON p.post_id = v.post_id
    	AND v.user_id = $1
	INNER JOIN users u
    	ON p.user_id = u.user_id
	LEFT JOIN follows f
    	ON f.following_id = p.user_id
    	AND f.follower_id = $1
	WHERE p.is_removed = FALSE AND p.post_id >$2 AND u.is_active = TRUE
	ORDER BY p.post_id DESC;`, userId, stopId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.PostResponse])

	return posts, err
}

func (r *PostRepository) GetManyOrderedByVotes(ctx context.Context, userId int, stopId int) ([]model.PostResponse, error) {
	rows, err := r.Db.Query(ctx, `
	SELECT
    	p.post_id, p.user_id, p.url, p.text, p.votes, 
    	u.username,
    	u.profile_picture_url,
    	CASE WHEN v.user_id IS NOT NULL THEN TRUE ELSE FALSE END AS voted,
		CASE WHEN f.follower_id IS NOT NULL THEN TRUE ELSE FALSE END AS following
	FROM posts p
	LEFT JOIN votes v
   		ON p.post_id = v.post_id
    	AND v.user_id = $1
	INNER JOIN users u
    ON p.user_id = u.user_id
	LEFT JOIN follows f
    	ON f.following_id = p.user_id
    	AND f.follower_id = $1
	WHERE p.is_removed = FALSE AND p.post_id >$2 AND u.is_active = TRUE
	ORDER BY p.votes DESC, p.post_id DESC;`, userId, stopId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.PostResponse])

	return posts, err
}

func (r *PostRepository) GetVote(ctx context.Context, postId int, userId int) (model.Vote, error) {

	var vote model.Vote

	err := r.Db.QueryRow(
		ctx,
		"SELECT post_id, user_id FROM votes WHERE post_id=$1 AND user_id=$2",
		postId, userId,
	).Scan(&vote.PostId, &vote.UserId)

	return vote, err
}

func (r *PostRepository) InsertVote(ctx context.Context, postId int, userId int) error {
	_, err := r.Db.Exec(
		ctx,
		"INSERT INTO votes (post_id, user_id) VALUES ($1, $2)",
		postId, userId,
	)

	if err != nil {
		return err
	}

	post, err := r.GetPost(ctx, postId)

	if err != nil {
		return err
	}

	_, err = r.Db.Exec(ctx, "UPDATE posts SET votes=$2 WHERE post_id=$1", post.PostId, post.Votes + 1)

	log.Println(err)

	return err
}

func (r *PostRepository) RemoveVote(ctx context.Context, postId int, userId int) error {
	tag, err := r.Db.Exec(
		ctx,
		"DELETE FROM votes WHERE post_id=$1 AND user_id=$2",
		postId, userId,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("0 columnas afectadas")
	}

	post, err := r.GetPost(ctx, postId)

	if err != nil {
		return err
	}

	_, err = r.Db.Exec(ctx, "UPDATE posts SET votes=$2 WHERE post_id=$1", post.PostId, post.Votes - 1)

	return err
}

func (r *PostRepository) GetPost(ctx context.Context, postId int) (model.Post, error) {

	rows, err := r.Db.Query(ctx,"SELECT post_id, user_id, url, votes, public_id, text FROM posts WHERE post_id=$1", postId)

	if err != nil {
		return model.Post{}, err
	}

	post, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[model.Post])

	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

func (r *PostRepository) RemovePost(ctx context.Context, postId int) error {
	tag, err := r.Db.Exec(
		ctx,
		"UPDATE posts SET is_removed=TRUE WHERE post_id=$1",
		postId,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("0 columnas afectadas")
	}

	return nil
}


func (r *PostRepository) GetManyOrderedByFavorites(ctx context.Context, userId int, stopId int) ([]model.PostResponse, error) {

	rows, err := r.Db.Query(ctx, `
	SELECT
    	p.post_id, p.user_id, p.url, p.text, p.votes, 
    	u.username,
    	u.profile_picture_url,
    	CASE WHEN v.user_id IS NOT NULL THEN TRUE ELSE FALSE END AS voted,
		CASE WHEN f.follower_id IS NOT NULL THEN TRUE ELSE FALSE END AS following
	FROM posts p
	INNER JOIN votes v
   		ON p.post_id = v.post_id
    	AND v.user_id = $1
	INNER JOIN users u
    	ON p.user_id = u.user_id
	LEFT JOIN follows f
    	ON f.following_id = p.user_id
    	AND f.follower_id = $1
	WHERE p.is_removed = FALSE AND p.post_id >$2 AND u.is_active = TRUE
	ORDER BY p.post_id DESC;`, userId, stopId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.PostResponse])

	return posts, err
}
