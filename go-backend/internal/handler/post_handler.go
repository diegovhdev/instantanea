package handler

import (
	"instantanea/internal/model"
	"instantanea/internal/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type PostHandler struct {
	Service    *service.PostService
	Validator  *validator.Validate
	middleware  Middleware
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
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
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
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	WriteJSON(w, http.StatusOK, posts)
}

func (h *PostHandler) GetPostsDispatcher(w http.ResponseWriter, r *http.Request) {

	orderedByStr := r.URL.Query().Get("ordered-by")
	var postLists []model.PostResponse
	var err error
	stopIdStr := r.URL.Query().Get("stopId")
	var stopId int;

	if stopIdStr == "" {
		stopId = -1
	} else {
		stopId, err = strconv.Atoi(stopIdStr)
		if err != nil {
			http.Error(w, "stopId tiene que ser un numero", http.StatusBadRequest)
			return
		}
	}

	switch orderedByStr {
	case "id":
		postLists, err = h.Service.GetPostsById(r.Context(), stopId)
	case "votes":
		postLists, err = h.Service.GetPostsByVotes(r.Context(), stopId)
	case "favorites":
	default:
		http.Error(w, "argumento de url invalido", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), GetErrorStatusCode(err))
		return
	}

	WriteJSON(w, http.StatusOK, postLists)
}

func (h *PostHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /posts", h.middleware.HandlerFunc(h.Post))
	mux.Handle("GET /posts", h.middleware.HandlerFunc(h.GetPostsDispatcher))
}

func NewPostHandler(service *service.PostService, validator *validator.Validate, middleware Middleware) PostHandler {
	return PostHandler{
		Service: service,
		Validator: validator,
		middleware: middleware,
	}
}