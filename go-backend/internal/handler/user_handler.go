package handler

import (
	"instantanea/internal/model"
	"instantanea/internal/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	Service *service.UserService
	Validator *validator.Validate
	middleware Middleware
}

func (h *UserHandler) PatchProfilePicture(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error al procesar formulario", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error al leer el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "el argumento del path tiene que ser un numero", http.StatusBadRequest)
		return
	}

	idContext, ok := r.Context().Value("UserID").(int)

	if !ok {
		http.Error(w, "error en el id de la sesion", http.StatusInternalServerError)
		return
	}

	if id != idContext {
		http.Error(w, "el id pedido y el que se tiene no coinciden", http.StatusForbidden)
		return
	}

	url, err := h.Service.UpdateProfilePicture(r.Context(), id, file)

	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	responseUpdateProfilePicture := model.ResponseUpdateProfilePicture{PictureURL: url}

	WriteJSON(w, http.StatusOK, responseUpdateProfilePicture)
}

func (h *UserHandler) PatchUsername(w http.ResponseWriter, r *http.Request) {
	requestUpdateUsername, err := ValidateJSON[model.RequestUpdateUsername](r.Body, h.Validator)

	if err != nil {
		http.Error(w, "las validaciones para los inputs fallaron", http.StatusBadRequest)
		return
	}

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "el argumento del path tiene que ser un numero", http.StatusBadRequest)
		return
	}

	idContext, ok := r.Context().Value("UserID").(int)

	if !ok {
		http.Error(w, "error en el id de la sesion", http.StatusInternalServerError)
		return
	}

	if id != idContext {
		http.Error(w, "el id pedido y el que se tiene no coinciden", http.StatusForbidden)
		return
	}

	err = h.Service.UpdateUsername(r.Context(), id, requestUpdateUsername)

	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) PatchEmail(w http.ResponseWriter, r *http.Request) {
	requestUpdateEmail, err := ValidateJSON[model.RequestUpdateEmail](r.Body, h.Validator)

	if err != nil {
		http.Error(w, "las validaciones para los inputs fallaron", http.StatusBadRequest)
		return
	}

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "el argumento del path tiene que ser un numero", http.StatusBadRequest)
		return
	}

	idContext, ok := r.Context().Value("UserID").(int)

	if !ok {
		http.Error(w, "error en el id de la sesion", http.StatusInternalServerError)
		return
	}

	if id != idContext {
		http.Error(w, "el id pedido y el que se tiene no coinciden", http.StatusForbidden)
		return
	}

	err = h.Service.UpdateEmail(r.Context(), id, requestUpdateEmail)


	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) PostPassword(w http.ResponseWriter, r *http.Request) {
	requestUpdatePassword, err := ValidateJSON[model.RequestUpdatePassword](r.Body, h.Validator)

	if err != nil {
		http.Error(w, "las validaciones para los inputs fallaron", http.StatusBadRequest)
		return
	}

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "el argumento del path tiene que ser un numero", http.StatusBadRequest)
		return
	}

	idContext, ok := r.Context().Value("UserID").(int)

	if !ok {
		http.Error(w, "error en el id de la sesion", http.StatusInternalServerError)
		return
	}

	if id != idContext {
		http.Error(w, "el id pedido y el que se tiene no coinciden", http.StatusForbidden)
		return
	}
	
	err = h.Service.UpdatePassword(r.Context(), id, requestUpdatePassword)


	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("PATCH /users/{id}/profile-picture", h.middleware.HandlerFunc(h.PatchProfilePicture))
	mux.Handle("PATCH /users/{id}/username", h.middleware.HandlerFunc(h.PatchUsername))
	mux.Handle("PATCH /users/{id}/email", h.middleware.HandlerFunc(h.PatchEmail))
	mux.Handle("POST /users/{id}/password", h.middleware.HandlerFunc(h.PostPassword))
}


func NewUserHanler(service *service.UserService, validator *validator.Validate, middleware Middleware) UserHandler {
	return UserHandler{
		Service:    service,
		Validator:  validator,
		middleware: middleware,
	}
}


