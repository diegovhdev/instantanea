package main

import (
	"context"
	"fmt"
	"instantanea/internal/middlewares"
	"instantanea/internal/users"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()


	connString := "postgresql://neondb_owner:npg_SNsXzW8JOf6o@ep-red-snow-aizqkaxu-pooler.c-4.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

	pool, err := pgxpool.New(ctx, connString)

	if err != nil {
		panic(err)
	}

	validate := validator.New()

	handleUser := users.Handler{
		Repository: &users.Repository{
			Db: pool,
		},
		Validator: validate,
	}

	mux := http.NewServeMux()

	handleUser.RegisterRoutes(mux)

	fmt.Println("Inicio el servidor")

	err = http.ListenAndServe(":8080", middlewares.CORS(mux))
	log.Fatal(err)
}