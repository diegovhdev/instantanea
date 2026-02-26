package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type SpaHandler struct {
	distDir string
	fs      http.Handler
}

func NewSpaHandler(dist string) *SpaHandler {
	return &SpaHandler{
		distDir: dist,
		fs:      http.FileServer(http.Dir(dist)),
	}
}

func (h *SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    path := filepath.Join(h.distDir, filepath.Clean(r.URL.Path))

    if !strings.HasPrefix(path, filepath.Clean(h.distDir)) {
        http.Error(w, "invalid path", http.StatusBadRequest)
        return
    }

    info, err := os.Stat(path)
    if err != nil || info.IsDir() {
        http.ServeFile(w, r, filepath.Join(h.distDir, "index.html"))
        return
    }

    h.fs.ServeHTTP(w, r)
}
