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

	id, code, err := VerifyId(r)

	if err != nil {
		http.Error(w, err.Error(), code)
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
	
	id, code, err := VerifyId(r)

	if err != nil {
		http.Error(w, err.Error(), code)
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

	id, code, err := VerifyId(r)

	if err != nil {
		http.Error(w, err.Error(), code)
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

	id, code, err := VerifyId(r)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	
	err = h.Service.UpdatePassword(r.Context(), id, requestUpdatePassword)


	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	requestDeleteUser, err := ValidateJSON[model.RequestDeleteUser](r.Body, h.Validator)

	if err != nil {
		http.Error(w, "las validaciones para los inputs fallaron", http.StatusBadRequest)
		return
	}

	id, code, err := VerifyId(r)

	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	
	err = h.Service.DeleteUser(r.Context(), id, requestDeleteUser)


	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


func (h *UserHandler) PostFollow(w http.ResponseWriter, r *http.Request) {
	userId, err := ExtractId(r) 

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	followingIdString := r.PathValue("id")
	followingId, err := strconv.Atoi(followingIdString)

	if err != nil {
		http.Error(w, "el id debe ser un numero", http.StatusBadRequest)
		return
	}

	err = h.Service.Follow(r.Context(), userId, followingId)

	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}



func (h *UserHandler) DeleteFollow(w http.ResponseWriter, r *http.Request) {
	userId, err := ExtractId(r) 

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	followingIdString := r.PathValue("id")
	followingId, err := strconv.Atoi(followingIdString)

	if err != nil {
		http.Error(w, "el id debe ser un numero", http.StatusBadRequest)
		return
	}

	err = h.Service.UnFollow(r.Context(), userId, followingId)

	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) GetFollowingUsers(w http.ResponseWriter, r *http.Request) {

	var users []model.UserFollowingResponse
	userIdString := r.PathValue("id")
	userId, err :=  strconv.Atoi(userIdString)

	if err != nil {
		http.Error(w, "el id debe ser un numero", http.StatusBadRequest)
		return
	}

	users, err = h.Service.GetFollowingUsers(r.Context(), userId)

	println(users)

	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	WriteJSON(w, 200, users)
}


func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("PATCH /users/{id}/profile-picture", h.middleware.HandlerFunc(h.PatchProfilePicture))
	mux.Handle("PATCH /users/{id}/username", h.middleware.HandlerFunc(h.PatchUsername))
	mux.Handle("PATCH /users/{id}/email", h.middleware.HandlerFunc(h.PatchEmail))
	mux.Handle("POST /users/{id}/password", h.middleware.HandlerFunc(h.PostPassword))
	mux.Handle("DELETE /users/{id}", h.middleware.HandlerFunc(h.DeleteUser))
	mux.Handle("POST /users/{id}/follow", h.middleware.HandlerFunc(h.PostFollow))
	mux.Handle("DELETE /users/{id}/follow", h.middleware.HandlerFunc(h.DeleteFollow))
	mux.Handle("GET /users/{id}/following", h.middleware.HandlerFunc(h.GetFollowingUsers))
}


func NewUserHanler(service *service.UserService, validator *validator.Validate, middleware Middleware) UserHandler {
	return UserHandler{
		Service:    service,
		Validator:  validator,
		middleware: middleware,
	}
}


