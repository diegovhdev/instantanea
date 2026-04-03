package service

import (
	"context"
	"instantanea/internal/model"
	"instantanea/internal/repository"
	"mime/multipart"
)

type PostService struct {
	Repository *repository.PostRepository
	uploader Uploader
}


func (s *PostService) Post(ctx context.Context, text string, file multipart.File) error {

	url, publicId, err := s.uploader.Upload(ctx, file)
	if err != nil {
		return err
	}

	userID, ok := ctx.Value("UserID").(int)
	if !ok {
		return &CustomError{nil, ErrUnauthenticatedUser}
	}

	post := model.Post{
		Text: text,
		UserId: userID,
		Url: url,
		PublicId: publicId,
	}

	if err := s.Repository.Insert(ctx, post); err != nil {
		return &CustomError{err, ErrInDatabase}
	}

	return nil
}

func (s *PostService) GetPosts(ctx context.Context, limit, offset int) ([]model.PostResponse, error) {
	posts, err := s.Repository.GetMany(ctx, limit, offset)
	if err != nil {
		return nil, &CustomError{err, ErrInDatabase}
	}
	if len(posts) == 0 {
		return posts, &CustomError{nil, ErrNoContent}
	}

	return posts, nil
}


func NewPostService(repository *repository.PostRepository, uploader Uploader) *PostService {
	return &PostService{
		Repository: repository,
		uploader: uploader,
	}
}