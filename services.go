package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type Lesson struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// (*GetLessonsHandler) ServeHTTP queries the db and populates the results in a slice of Lesson.
func (h *GetLessonsHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	db, err := gorm.Open(postgres.Open(CreateConnString()))

	if err != nil {
		h.log.Error("the db connection failed")
	}

	var lessons []Lesson
	db.Select([]string{"id", "title", "description"}).Find(&lessons)

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(lessons); err != nil {
		h.log.Error("something went wrong serializing query results")
	}
}

// (*GetLessonByIdHandler) ServeHTTP queries the db for a Lesson matching the id.
func (h *GetLessonByIdHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(CreateConnString()))
	if err != nil {
		h.log.Error("the db connection failed")
	}

	var lesson Lesson
	values := mux.Vars(r)

	db.First(&lesson, values["id"])

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(lesson); err != nil {
		h.log.Error("error serializing lesson data to json")
	}
}

// (*GenerateLessonHandler) ServeHTTP deserializes a new Lesson from the request body and returns it to the caller.
func (h *GenerateLessonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
		h.log.Error("Wrong HTTP Method used", r.Method)
		return
	}

	var lesson Lesson
	if err := json.NewDecoder(r.Body).Decode(&lesson); err != nil {
		http.Error(w, "Error deserializing request body.", http.StatusBadRequest)
		h.log.Error("Encountered error deserializing request", err)
		return
	}

	_, err := fmt.Fprintf(w,
		"Received lesson of Id: %d, Title: %s, Description: %s",
		lesson.Id, lesson.Title, lesson.Description)
	if err != nil {
		h.log.Warn("fmt err: %v", err)
	}
}
