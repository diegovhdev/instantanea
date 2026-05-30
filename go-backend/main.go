package main

import (
	"context"
	"embed"
	"fmt"
	"instantanea/internal/handler"
	"instantanea/internal/middleware"
	"instantanea/internal/repository"
	"instantanea/internal/service"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

//go:embed all:build
var buildFS embed.FS

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

	uploader := service.NewCloudinaryUploader(cdl)

	authHandler := handler.NewAuthHanler(
		service.NewAuthService(&repository.UserRepository{Db: pool}), 
		validate,
		middleware.Auth,
	)

	postHandler := handler.NewPostHandler(
		service.NewPostService(&repository.PostRepository{Db: pool}, uploader), 
		validate,
		middleware.Auth,
	)

	userHandler := handler.NewUserHanler(
		service.NewUserService(&repository.UserRepository{Db: pool}, uploader),
		validate,
		middleware.Auth,
	)

	mux := http.NewServeMux()

	authHandler.RegisterRoutes(mux)
	postHandler.RegisterRoutes(mux)
	userHandler.RegisterRoutes(mux)

    stripped, err := fs.Sub(buildFS, "build")
    if err != nil {
        log.Fatal(err)
    }
    mux.Handle("/", NewSpaHandler(stripped))

	fmt.Println("Inicio el servidor en el puerto :8080")

	wrappper := middleware.Logger(middleware.CORS(mux))

	err = http.ListenAndServe(":8080", wrappper)
	log.Fatal(err)
}