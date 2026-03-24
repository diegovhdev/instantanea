package main

import (
	"context"
	"fmt"
	"instantanea/internal/handler"
	"instantanea/internal/middleware"
	"instantanea/internal/repository"
	"instantanea/internal/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
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

	userHandler := handler.NewUserHanler(service.NewUserService(&repository.UserRepository{Db: pool}), validate)

	cld, _ := cloudinary.NewFromParams(
    	os.Getenv("CLOUDINARY_CLOUD_NAME"),
    	os.Getenv("CLOUDINARY_API_KEY"),
    	os.Getenv("CLOUDINARY_API_SECRET"),
	)

	postHandler := handler.NewPostHandler(service.NewPostService(&repository.PostRepository{Db: pool}, cld), validate)

	mux := http.NewServeMux()

	userHandler.RegisterRoutes(mux)
	postHandler.RegisterRoutes(mux)

	fmt.Println("Inicio el servidor")

	err = http.ListenAndServe(":8080", middleware.CORS(mux))
	log.Fatal(err)
}