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


	connString := "placeholder string de conexión"

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