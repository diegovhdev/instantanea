package main

import (
	"io/fs"
	"net/http"
	"strings"
)

type SpaHandler struct {
    fs http.FileSystem
}

func NewSpaHandler(fsys fs.FS) *SpaHandler {
    return &SpaHandler{fs: http.FS(fsys)}
}

func (s *SpaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fileServer := http.FileServer(s.fs)

    // Limpiar el path y quitar el "/" inicial para abrir con fs
    path := strings.TrimPrefix(r.URL.Path, "/")
    if path == "" {
        path = "index.html"
    }

    f, err := s.fs.Open(path)
    if err != nil {
        // Archivo no existe → fallback al SPA
        r2 := r.Clone(r.Context())
        r2.URL.Path = "/"
        fileServer.ServeHTTP(w, r2)
        return
    }
    defer f.Close()

    // Si es directorio, servir index.html
    stat, err := f.Stat()
    if err != nil || stat.IsDir() {
        r2 := r.Clone(r.Context())
        r2.URL.Path = "/"
        fileServer.ServeHTTP(w, r2)
        return
    }

    fileServer.ServeHTTP(w, r)
}