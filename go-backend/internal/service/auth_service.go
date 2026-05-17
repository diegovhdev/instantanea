package service

import (
	"context"
	"instantanea/internal/model"
	"instantanea/internal/repository"
	"strconv"
)


type AuthService struct {
	Repository *repository.UserRepository
}

func (s *AuthService) Register(ctx context.Context, user model.UserRegisterRequest) error {

	if _, err := s.Repository.FindByUsername(ctx, user.Username); err == nil {
		return &CustomError{nil, ErrUsernameAlreadyExists}
	}

	if _, err := s.Repository.FindByEmail(ctx, user.Email); err == nil {
		return &CustomError{nil, ErrEmailAlreadyExists}
	}

	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return &CustomError{err, ErrHashingPassword}
	}

	user.Password = hashedPassword

	if err := s.Repository.Insert(ctx, user.ToUser()); err != nil {
		return &CustomError{err, ErrInternal}
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, user model.UserLoginRequest) (string, model.UserResponse, error) {

	userFound, err := s.Repository.FindByUsername(ctx, user.Username)

	if err != nil || !userFound.IsActive {
		return "", model.UserResponse{}, &CustomError{err, ErrUserNotFound}
	}

	if err := CheckPassword(userFound.Password, user.Password); err != nil {
		return "", model.UserResponse{}, &CustomError{err, ErrIncorrectPassword}
	}

	token, err := GenerateToken(strconv.Itoa(userFound.UserId))

	if err != nil {
		return "", model.UserResponse{}, &CustomError{err, ErrTokenGeneration}
	}

	return token, userFound.ToUserResponse(), nil
}

func NewAuthService(repository *repository.UserRepository) *AuthService {
	return &AuthService{
		Repository: repository,
	}
}