package main

import "log/slog"

// GetLessonsHandler is an http.Handler for getting all Lesson from the db.
type GetLessonsHandler struct {
	log *slog.Logger
}

// GetLessonByIdHandler is an http.Handler for getting a Lesson from the db.
type GetLessonByIdHandler struct {
	log *slog.Logger
}

// GenerateLessonHandler is an http.Handler for generating a Lesson from http.Request.
type GenerateLessonHandler struct {
	log *slog.Logger
}

// -----

// NewGetLessonsHandler builds a GetLessonsHandler.
func NewGetLessonsHandler(l *slog.Logger) *GetLessonsHandler {
	return &GetLessonsHandler{log: l}
}

// NewGetLessonByIdHandler builds a GetLessonByIdHandler.
func NewGetLessonByIdHandler(l *slog.Logger) *GetLessonByIdHandler {
	return &GetLessonByIdHandler{log: l}
}

// NewGenerateLessonHandler builds a GenerateLessonHandler.
func NewGenerateLessonHandler(l *slog.Logger) *GenerateLessonHandler {
	return &GenerateLessonHandler{log: l}
}
