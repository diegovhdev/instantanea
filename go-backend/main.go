package main

import (
	"context"
	"fmt"
	"instantanea/internal/middlewares"
	"instantanea/internal/posts"
	"instantanea/internal/users"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	
	 if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()


	connString := os.Getenv("DATABASE_STRING")

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

	handlePost := posts.Handler{
		Repository: &posts.Repository{
			Db: pool,
		},
		Validator: validate,
	}

	mux := http.NewServeMux()

	handleUser.RegisterRoutes(mux)
	handlePost.RegisterRoutes(mux)

	fmt.Println("Inicio el servidor")

	err = http.ListenAndServe(":8080", middlewares.CORS(mux))
	log.Fatal(err)
}