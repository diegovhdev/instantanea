package handler

import (
	"instantanea/internal/model"
	"instantanea/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	Service   *service.AuthService
	Validator *validator.Validate
	middleware Middleware
}


func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {


	user, err := ValidateJSON[model.UserRegisterRequest](r.Body, h.Validator)

	if err != nil {
		http.Error(w, "las validaciones para los inputs fallaron", http.StatusBadRequest)
		return
	}

	err = h.Service.Register(r.Context(), user)

	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusCreated)
}


func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	user, err := ValidateJSON[model.UserLoginRequest](r.Body, h.Validator)

	if err != nil {
		http.Error(w, "las validaciones para los inputs fallaron", http.StatusBadRequest)
		return
	}

	token, err := h.Service.Login(r.Context(), user)

	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	SetCookie(w, "access_token", token)

	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	RemoveCookie(w, "access_token")
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Test(w http.ResponseWriter, r *http.Request) {
	a := struct{
		Message string `json:"message"`
	}{"Aceddiste a la ruta protegida"}

	WriteJSON(w, http.StatusOK, a)
}

func (h *AuthHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /auth/register", http.HandlerFunc(h.Register))
	mux.Handle("POST /auth/login", http.HandlerFunc(h.Login))
	mux.Handle("GET /auth/logout", http.HandlerFunc(h.Logout))
	mux.Handle("GET /test", h.middleware.HandlerFunc(h.Test))
}

func NewAuthHanler(service *service.AuthService, validator *validator.Validate, middleware Middleware) AuthHandler {
	return AuthHandler{
		Service: service,
		Validator: validator,
		middleware: middleware,
	}
}