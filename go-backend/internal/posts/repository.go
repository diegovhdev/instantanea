package posts

import (
	"context"

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