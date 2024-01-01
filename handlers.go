package main

// GetLessonsHandler is an http.Handler for getting all Lesson from the db.
type GetLessonsHandler struct{}

// NewGetLessonsHandler builds a GetLessonsHandler.
func NewGetLessonsHandler() *GetLessonsHandler {
	return &GetLessonsHandler{}
}

// GetLessonByIdHandler is an http.Handler for getting a Lesson from the db.
type GetLessonByIdHandler struct{}

// NewGetLessonByIdHandler builds a GetLessonByIdHandler.
func NewGetLessonByIdHandler() *GetLessonByIdHandler {
	return &GetLessonByIdHandler{}
}

// GenerateLessonHandler is an http.Handler for generating a Lesson from http.Request.
type GenerateLessonHandler struct{}

// NewGenerateLessonHandler builds a GenerateLessonHandler.
func NewGenerateLessonHandler() *GenerateLessonHandler {
	return &GenerateLessonHandler{}
}
