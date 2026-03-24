package handler

import (
	"instantanea/internal/middleware"
	"instantanea/internal/model"
	"instantanea/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	Service   *service.UserService
	Validator *validator.Validate
}


func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {

	user, err := ValidateJSON[model.User](r.Body, h.Validator)

	if err != nil {
		http.Error(w, "las validaciones para los inputs fallaron", http.StatusBadRequest)
		return
	}

	err = h.Service.Register(r.Context(), user)

	if err != nil {
		switch err {
		case service.ErrUsernameAlreadyExists, service.ErrUsernameAlreadyExists:
			http.Error(w, err.Error(), http.StatusConflict)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}


func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {

	user, err := ValidateJSON[model.UserRequest](r.Body, h.Validator)

	if err != nil {
		http.Error(w, "las validaciones para los inputs fallaron", http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(r.Context(), user)

	if err != nil {
		switch err {
		case service.ErrUserNotFound:
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		case service.ErrIncorrectPassword:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	http.SetCookie(w, &http.Cookie{
		Name: "access_token",
		Value: token,
		HttpOnly: true,
		Path: "/",
		SameSite: http.SameSiteLaxMode,
		Secure: true,
	})

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "access_token",
		Value: "",
		Path: "/",
		MaxAge: -1,
		Secure: true,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})
}

func (h *UserHandler) Test(w http.ResponseWriter, r *http.Request) {
	a := struct{
		Message string `json:"message"`
	}{"Aceddiste a la ruta protegida"}

	WriteJSON(w, http.StatusOK, a)
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /register", http.HandlerFunc(h.Register))
	mux.Handle("POST /login", http.HandlerFunc(h.Login))
	mux.Handle("GET /logout", middleware.Auth(h.Logout))
	mux.Handle("GET /test", middleware.Auth(h.Test))
}

func NewUserHanler(service *service.UserService, validator *validator.Validate) UserHandler {
	return UserHandler{
		Service: service,
		Validator: validator,
	}
}