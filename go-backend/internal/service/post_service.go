package service

import (
	"context"
	"errors"
	"instantanea/internal/model"
	"instantanea/internal/repository"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var ErrUploadingImage = errors.New("error al subir la imagen")
var ErrUnauthenticatedUser = errors.New("Usuario no autentificado")
var ErrInDatabase = errors.New("Hubo un error con la base de datos")
var ErrNoContent = errors.New("No hay más contenido")

type PostService struct {
	Repository *repository.PostRepository
	Cloudinary *cloudinary.Cloudinary
}


func (s *PostService) Post(ctx context.Context, text string, file multipart.File) error {

	result, err := s.Cloudinary.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		return ErrUploadingImage
	}

	userID, ok := ctx.Value("UserID").(int)
	if !ok {
		return ErrUnauthenticatedUser
	}

	post := model.Post{
		Text: text,
		UserId: userID,
		Url: result.SecureURL,
		PublicId: result.PublicID,
	}

	if err := s.Repository.Insert(ctx, post); err != nil {
		return ErrInDatabase
	}

	return nil
}

func (s *PostService) GetPosts(ctx context.Context, limit, offset int) ([]model.PostResponse, error) {
	posts, err := s.Repository.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, ErrInDatabase
	}
	if len(posts) == 0 {
		return posts, ErrNoContent
	}

	return posts, nil
}


func NewPostService(repository *repository.PostRepository, clodinary *cloudinary.Cloudinary) *PostService {
	return &PostService{
		Repository: repository,
		Cloudinary: clodinary,
	}
}