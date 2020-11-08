package http

import (
	"log"
	"net/http"
	"os"
)

func middlewareJsonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func middlewareIndexLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetOutput(os.Stdout)
		log.Println(r.Method, r.URL)

		next.ServeHTTP(w, r)
	})
}
