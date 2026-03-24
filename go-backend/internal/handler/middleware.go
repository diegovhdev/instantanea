package handler

import "net/http"

type Middleware func(next http.Handler) http.Handler

func (m Middleware) HandlerFunc(next http.HandlerFunc) http.Handler {
	return m(next)
}
