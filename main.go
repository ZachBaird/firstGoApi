package main

import (
	"go.uber.org/fx"
	"log/slog"
	"net/http"
	"os"
)

func NewLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func main() {
	app := fx.New(
		fx.Provide(NewHttpServer),
		fx.Provide(NewServeMux),
		fx.Provide(NewGetLessonsHandler),
		fx.Provide(NewGetLessonByIdHandler),
		fx.Provide(NewGenerateLessonHandler),
		fx.Provide(NewLogger),
		fx.Invoke(func(*http.Server) {}),
	)

	app.Run()
}
