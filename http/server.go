package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type domain interface {
	GetIsDisposable(string) bool
}

type Server struct {
	domain domain
}

type ErrorResponse struct {
	Detail string `json:"detail"`
}

type Response struct {
	IsDisposable bool `json:"isDisposable"`
}

func NewServer(d domain) Server {
	return Server{
		domain: d,
	}
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	if email == "" {
		errorResponse, _ := json.Marshal(ErrorResponse{
			Detail: "email query parameter should be set.",
		})

		w.WriteHeader(400)
		w.Write(errorResponse)

		return
	}

	parts := strings.Split(email, "@")
	emailDomain := parts[len(parts)-1]

	isDisposable := s.domain.GetIsDisposable(emailDomain)

	jsonBytes, _ := json.Marshal(Response{
		IsDisposable: isDisposable,
	})

	w.Write(jsonBytes)
}

func (s *Server) Run() {
	mux := http.NewServeMux()

	mux.Handle("/", middlewareIndexLogger(middlewareJsonHeader(http.HandlerFunc(s.index))))

	log.Println("Starting to listen on 0.0.0.0:80...")
	log.Fatal(http.ListenAndServe(":80", mux))
}
