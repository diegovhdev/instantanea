package service

import (
	"context"
	"instantanea/internal/model"
	"instantanea/internal/repository"
	"mime/multipart"
)

type UserService struct {
	Repository *repository.UserRepository
	uploader Uploader
}



func (s *UserService) UpdateUsername(ctx context.Context, userId int, data model.RequestUpdateUsername) (error) {

	if _, err := s.Repository.Find(ctx, userId); err != nil {
		return &CustomError{err, ErrUserNotFound}
	}

	if _, err := s.Repository.FindByUsername(ctx, data.Username); err == nil {
		return &CustomError{err, ErrUsernameAlreadyExists}
	}

	err := s.Repository.UpdateUsername(ctx, userId, data.Username)

	if err != nil {
		return &CustomError{err, ErrUsernameCouldNotBeUpdated}
	}

	return nil
}

func (s *UserService) UpdateEmail(ctx context.Context, userId int, data model.RequestUpdateEmail) (error) {

	user, err := s.Repository.Find(ctx, userId)

	if err != nil {
		return &CustomError{err, ErrUserNotFound}
	}

	if user.Email != data.Email {
		return &CustomError{nil, ErrIncorrectEmail}
	}

	if _, err := s.Repository.FindByEmail(ctx, data.NewEmail); err == nil {
		return &CustomError{err, ErrEmailAlreadyExists}
	}

	err = s.Repository.UpdateEmail(ctx, userId, data.NewEmail)

	if err != nil {
		return &CustomError{err, ErrEmailCouldNotBeUpdated}
	}

	return nil
}

func (s *UserService) UpdatePassword(ctx context.Context, userId int, data model.RequestUpdatePassword) (error) {

	user, err := s.Repository.Find(ctx, userId)

	if err != nil {
		return &CustomError{err, ErrUserNotFound}
	}

	if CheckPassword(user.Password, data.Password) != nil {
		return &CustomError{nil, ErrIncorrectPassword}
	}

	newPassword, err := HashPassword(data.NewPassword)

	if err != nil {
		return &CustomError{err, ErrHashingPassword}
	}

	err = s.Repository.UpdatePassword(ctx, userId, newPassword)

	if err != nil {
		return &CustomError{err, ErrPasswordCouldNotBeUpdate}
	}

	return nil
}

func (s *UserService) UpdateProfilePicture(ctx context.Context, userId int, file multipart.File) (string, error) {

	url, _, err := s.uploader.Upload(ctx, file)
	if err != nil {
		return "", err
	}

	_, err = s.Repository.Find(ctx, userId)

	if err != nil {
		return "", &CustomError{err, ErrUserNotFound}
	}

	err = s.Repository.UpdateProfilePicture(ctx, userId, url)

	if err != nil {
		return "", &CustomError{err, ErrProfilePictureCouldNotBeUpdate}
	}

	return url, nil
}

func NewUserService(repository *repository.UserRepository, uploader Uploader) *UserService {
	return &UserService{
		Repository: repository,
		uploader: uploader,
	}
}