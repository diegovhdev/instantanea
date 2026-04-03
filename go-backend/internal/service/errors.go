package service

import (
	"errors"
	"log"
)


var ErrEmailAlreadyExists = errors.New("el correo ya existe")
var ErrUsernameAlreadyExists = errors.New("el usuario ya existe")
var ErrHashingPassword = errors.New("error en hashear la contraseña")
var ErrInternal = errors.New("error desconocido")
var ErrUserNotFound = errors.New("usuario no encontrado")
var ErrIncorrectPassword = errors.New("contraseña incorrecta")
var ErrTokenGeneration = errors.New("error en la generación del token")
var ErrUploadingImage = errors.New("error al subir la imagen")
var ErrUnauthenticatedUser = errors.New("Usuario no autentificado")
var ErrInDatabase = errors.New("Hubo un error con la base de datos")
var ErrNoContent = errors.New("No hay más contenido")


type CustomError struct {
	InternalError error
	PublicError   error
}

func (e *CustomError) Error() string {
	println("\n")
	if e.InternalError != nil {
		log.Println(e.InternalError)
	}
	log.Println(e.PublicError)
	return e.PublicError.Error()
}