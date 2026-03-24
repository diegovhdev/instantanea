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
	
	cdl, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	if err != nil {
		panic(err)
	}

	validate := validator.New()

	userHandler := handler.NewAuthHanler(
		service.NewAuthService(&repository.UserRepository{Db: pool}), 
		validate,
		middleware.Auth,
	)

	postHandler := handler.NewPostHandler(
		service.NewPostService(&repository.PostRepository{Db: pool}, service.NewCloudinaryUploader(cdl)), 
		validate,
		middleware.Auth,
	)

	mux := http.NewServeMux()

	userHandler.RegisterRoutes(mux)
	postHandler.RegisterRoutes(mux)

	fmt.Println("Inicio el servidor en el puerto :8080")

	err = http.ListenAndServe(":8080", middleware.CORS(mux))
	log.Fatal(err)
}