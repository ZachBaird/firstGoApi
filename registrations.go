package main

import (
	"context"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log/slog"
	"net"
	"net/http"
)

// NewHttpServer registers an http.Server with the application.
func NewHttpServer(lc fx.Lifecycle, r *mux.Router, l *slog.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: r}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			l.Info("Starting server at ", srv.Addr)
			go srv.Serve(listener)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

// NewServeMux builds a ServeMux that routes requests to the given handlers.
func NewServeMux(getLessonsHandler *GetLessonsHandler,
	getLessonByIdHandler *GetLessonByIdHandler,
	generateLessonHandler *GenerateLessonHandler) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/lessons", getLessonsHandler).Methods("GET")
	r.Handle("/lessons/{id}", getLessonByIdHandler).Methods("GET")
	r.Handle("/lessons", generateLessonHandler).Methods("POST")
	return r
}
