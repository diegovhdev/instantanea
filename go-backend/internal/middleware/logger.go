package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
    http.ResponseWriter
    StatusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(statusCode int) {
    lrw.StatusCode = statusCode
    lrw.ResponseWriter.WriteHeader(statusCode)
}


func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := &loggingResponseWriter{
			ResponseWriter: w,
			StatusCode: http.StatusOK,
		}

		next.ServeHTTP(lrw, r)
		fmt.Println()
		log.Printf(
    		"method=%s path=%s status=%d latency=%s",
    		r.Method,
    		r.URL.Path,
    		lrw.StatusCode,
			time.Since(start),
		)
	})
}