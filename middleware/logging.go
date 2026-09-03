package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)

		log.Printf("[%s] %s took %v", r.Method, r.URL.Path, time.Since(start))
	}
}
