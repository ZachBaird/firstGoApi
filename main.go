package main

import (
	"go.uber.org/fx"
	"net/http"
)

func main() {
	app := fx.New(
		fx.Provide(NewHttpServer),
		fx.Provide(NewServeMux),
		fx.Provide(NewGetLessonsHandler),
		fx.Provide(NewGetLessonByIdHandler),
		fx.Provide(NewGenerateLessonHandler),
		fx.Invoke(func(*http.Server) {}),
	)

	app.Run()
}
