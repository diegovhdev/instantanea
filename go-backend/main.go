package main

import (
	"context"
	"fmt"
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


	connString := "este string tiene que ser remplazado con el string de conexion a la base de datos"

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

	err = http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}