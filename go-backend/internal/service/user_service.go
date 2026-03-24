package service

import (
	"context"
	"errors"
	"instantanea/internal/model"
	"instantanea/internal/repository"
	"strconv"
)

var ErrEmailAlreadyExists = errors.New("el correo ya existe")
var ErrUsernameAlreadyExists = errors.New("el usuario ya existe")
var ErrHashingPassword = errors.New("error en hashear la contraseña")
var ErrInternal = errors.New("error desconocido")
var ErrUserNotFound = errors.New("usuario no encontrado")
var ErrIncorrectPassword = errors.New("contraseña incorrecta")
var ErrTokenGeneration = errors.New("error en la generación del token")

type UserService struct {
	Repository *repository.UserRepository
}

func (s *UserService) Register(ctx context.Context, user model.User) error {

	if _, err := s.Repository.FindByUsername(ctx, user.Username); err == nil {
		return ErrUsernameAlreadyExists
	}

	if _, err := s.Repository.FindByEmail(ctx, user.Email); err == nil {
		return ErrEmailAlreadyExists
	}

	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return ErrHashingPassword
	}

	user.Password = hashedPassword

	if err := s.Repository.Insert(ctx, user); err != nil {
		return ErrInternal
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, user model.UserRequest) (string, error) {

	userFound, err := s.Repository.FindByUsername(ctx, user.Username)

	if err != nil {
		return "", ErrUserNotFound
	}

	if err := CheckPassword(userFound.Password, user.Password); err != nil {
		return "", ErrIncorrectPassword
	}

	token, err := GenerateToken(strconv.Itoa(userFound.UserId))

	if err != nil {
		return "", ErrTokenGeneration
	}

	return token, nil
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{
		Repository: repository,
	}
}