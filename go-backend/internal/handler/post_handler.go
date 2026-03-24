package handler

import (
	"encoding/json"
	"instantanea/internal/middleware"
	"instantanea/internal/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type PostHandler struct {
	Service    *service.PostService
	Validator  *validator.Validate
}

func (h *PostHandler) Post(w http.ResponseWriter, r *http.Request) {


	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error al procesar formulario", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error al leer el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	err = h.Service.Post(r.Context(), text, file)

	if err != nil {
		switch err {
		case service.ErrUploadingImage, service.ErrUploadingImage:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		case service.ErrUnauthenticatedUser:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *PostHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	if limitStr == "" {
		http.Error(w, "el parametro limit esta vacio", http.StatusBadRequest)
		return
	}
	if offsetStr == "" {
		http.Error(w, "el parametro offset esta vacio", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 0 {
		http.Error(w, "limit debe ser un numero entero positivo", http.StatusBadRequest)
		return
	}
	offset, err := strconv.Atoi(offsetStr) 
	if err != nil || offset < 0 {
		http.Error(w, "limit debe ser un numero entero positivo", http.StatusBadRequest)
		return
	}

	posts, err := h.Service.GetPosts(r.Context(), limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(posts) == 0 {
		http.Error(w, "No hay más contenido", http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}


func (h *PostHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /posts", middleware.Auth(h.Post))
	mux.Handle("GET /posts", middleware.Auth(h.ListPosts))
}

func NewPostHandler(service *service.PostService, validator *validator.Validate) PostHandler {
	return PostHandler{
		Service: service,
		Validator: validator,
	}
}