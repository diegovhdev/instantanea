package posts

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Db *pgxpool.Pool
}

func (r *Repository) Insert(post Post, ctx context.Context) (error) {
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

func (r *Repository) GetMany(limit int, offset int, ctx context.Context) ([]PostResponse, error) {

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

	posts, err := pgx.CollectRows(rows, pgx.RowToStructByName[PostResponse])

	return posts, err
}