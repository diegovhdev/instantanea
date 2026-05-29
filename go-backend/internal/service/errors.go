package service

import (
	"errors"
	"log"
)


var ErrEmailAlreadyExists = errors.New("correo ya en uso")
var ErrUsernameAlreadyExists = errors.New("usuaruo ya en uso")
var ErrHashingPassword = errors.New("error en hashear la contraseña")
var ErrInternal = errors.New("error desconocido")
var ErrUserNotFound = errors.New("usuario no existe")
var ErrIncorrectPassword = errors.New("contraseña incorrecta")
var ErrTokenGeneration = errors.New("error en la generación del token")
var ErrUploadingImage = errors.New("error al subir la imagen")
var ErrUnauthenticatedUser = errors.New("Usuario no autentificado")
var ErrInDatabase = errors.New("Hubo un error con la base de datos")
var ErrNoContent = errors.New("No hay más contenido")
var ErrIncorrectEmail = errors.New("correo incorrecto")
var ErrUsernameCouldNotBeUpdated = errors.New("no se pudo actualizar el correo de usuario")
var ErrEmailCouldNotBeUpdated = errors.New("no se pudo actualizar el correo")
var ErrPasswordCouldNotBeUpdate = errors.New("no se pudo actualizar la contraseña")
var ErrProfilePictureCouldNotBeUpdate = errors.New("no se pudo actualizar la foto de perfil")
var ErrUserCouldNotBeDeleted = errors.New("No se pudo eliminar el usuario")
var ErrAlreadyVoted = errors.New("Ya hay un voto registrado")
var ErrAlreadyFollowing = errors.New("Ya se esta siguiendo al usuario")
var ErrAlreadyNotFollowing = errors.New("No se esta siguiendo al usuario")


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