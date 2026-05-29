package handler

import (
	"instantanea/internal/service"
	"net/http"
)


func GetErrorStatusCode(err error) int {
	customError, ok := err.(*service.CustomError)

	if !ok {
		return http.StatusInternalServerError
	}

	switch customError.PublicError {
		case service.ErrUsernameAlreadyExists, service.ErrEmailAlreadyExists:
			return http.StatusConflict
		case service.ErrUserNotFound:
			return http.StatusNotFound
		case service.ErrIncorrectPassword, service.ErrUnauthenticatedUser:
			return http.StatusUnauthorized
		case service.ErrNoContent:
			return http.StatusNoContent
		case service.ErrAlreadyVoted, service.ErrAlreadyFollowing, service.ErrAlreadyNotFollowing:
			return http.StatusBadRequest
		default:
			return http.StatusInternalServerError
	}
}