package users

import (
	"encoding/json"
	"instantanea/internal/helpers"
	"instantanea/internal/middlewares"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	*Repository
	Validator *validator.Validate
}


func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var user User


	if err := json.NewDecoder(r.Body).Decode(&user);  err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := helpers.ValidateStruct(h.Validator, user); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		http.Error(w, "ERROR IN HASHING", http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	if err := h.Insert(user, r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "ID INVALIDO", http.StatusBadRequest)
		return
	}

	user, err := h.Find(id, r.Context())
	if err != nil {
		http.Error(w, "USER NOT FOUND", http.StatusNotFound)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, user)
}


func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var user UserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := helpers.ValidateStruct(h.Validator, user); err != nil {
		helpers.WriteJSON(w, http.StatusBadRequest, err)
		return
	}

	userFound, err := h.FindByUsername(user.Username, r.Context())

	if err != nil {
		http.Error(w, "USER NOT FOUND", http.StatusNotFound)
		return
	}

	if err := helpers.CheckPassword(userFound.Password, user.Password); err != nil {
		http.Error(w, "INCORRECT PASSWORD", http.StatusUnauthorized)
		return
	}

	token, err := helpers.GenerateToken(userFound.Id)

	if err != nil {
		http.Error(w, "TOKEN GENERATION ERROR", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "access_token",
		Value: token,
		HttpOnly: true,
		Secure: true,
		Path: "/",
	})
}

func (h *Handler) Test(w http.ResponseWriter, r *http.Request) {
	a := struct{
		Message string `json:"message"`
	}{"Aceddiste a la ruta protegida"}

	helpers.WriteJSON(w, http.StatusOK, a)
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /register", http.HandlerFunc(h.Register))
	mux.Handle("POST /login", http.HandlerFunc(h.Login))
	mux.Handle("GET /test", middlewares.Auth(http.HandlerFunc(h.Test)))
}