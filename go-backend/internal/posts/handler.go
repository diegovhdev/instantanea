package posts

import (
	"context"
	"instantanea/internal/middlewares"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	*Repository
	Validator *validator.Validate
}

func (h *Handler) Upload(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(10 << 20)


    file, _, err := r.FormFile("image") // "image" = nombre del campo del frontend
    if err != nil {
        http.Error(w, "Error al leer el archivo", http.StatusBadRequest)
        return
    }
    defer file.Close()


    cld, err := cloudinary.NewFromParams(
        os.Getenv("CLOUDINARY_CLOUD_NAME"),
        os.Getenv("CLOUDINARY_API_KEY"),
        os.Getenv("CLOUDINARY_API_SECRET"),
    )
    if err != nil {
        http.Error(w, "Error al conectar con Cloudinary", http.StatusInternalServerError)
        return
    }

    ctx := context.Background()
    result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
    })

    if err != nil {
        http.Error(w, "Error al subir a Cloudinary", http.StatusInternalServerError)
        return
    }

	post := Post{
		Text: r.FormValue("text"),
		UserId: r.Context().Value("UserID").(int),
		Url: result.SecureURL,
		PublicId: result.PublicID,
	}

	if err := h.Insert(post, r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
}


func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /upload", middlewares.Auth(http.HandlerFunc(h.Upload)))
}