package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func New(logger *log.Logger) *Server {
	handler := http.NewServeMux()
	handler.Handle("GET /", http.HandlerFunc(handlers.Main))
	handler.Handle("POST /upload", http.HandlerFunc(handlers.Upload))

	return &Server{
		Logger: logger,
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      handler,
			ReadTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 10,
			IdleTimeout:  time.Second * 15,
			ErrorLog:     logger,
		},
	}
}
