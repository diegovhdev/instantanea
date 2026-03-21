package posts

import (
	"encoding/json"
	"instantanea/internal/middlewares"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	*Repository
	Validator  *validator.Validate
    Cloudinary *cloudinary.Cloudinary
}

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {

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

	result, err := h.Cloudinary.Upload.Upload(r.Context(), file, uploader.UploadParams{})
	if err != nil {
		http.Error(w, "Error al subir a Cloudinary", http.StatusInternalServerError)
		return
	}

	userID, ok := r.Context().Value("UserID").(int)
	if !ok {
		http.Error(w, "Usuario no autenticado", http.StatusUnauthorized)
		return
	}

	post := Post{
		Text: r.FormValue("text"),
		UserId: userID,
		Url: result.SecureURL,
		PublicId: result.PublicID,
	}

	if err := h.Insert(post, r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Feed(w http.ResponseWriter, r *http.Request) {
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

	posts, err := h.GetMany(limit, offset, r.Context())
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


func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /upload", middlewares.Auth(http.HandlerFunc(h.Upload)))
	mux.Handle("GET /feed", http.HandlerFunc(h.Feed))
}