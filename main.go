package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

const dbName string = "theodinproject_development"
const port string = "5432"
const sslMode string = "disable"

type Lesson struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// createConnString combines package-level variables and environment vars to build a postgresql conn string.
func createConnString() string {
	user := os.Getenv("DATABASE_USER")
	pw := os.Getenv("DATABASE_PASSWORD")
	return fmt.Sprintf("user=%s port=%s password=%s dbname=%s sslmode=%s", user, port, pw, dbName, sslMode)
}

// getLessons queries the db and populates the results in a slice of Lesson.
func getLessons(w http.ResponseWriter, r *http.Request) {
	connString := createConnString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("the db connection failed")
	}

	rows, err := db.Query("SELECT id, title, description FROM lessons")
	if err != nil {
		log.Fatal("failed to query db")
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		log.Fatal("failed to get columns")
	}

	for _, col := range cols {
		fmt.Fprintf(w, "%s  ", col)
	}
	fmt.Fprintf(w, "\n\n")

	var lessons []Lesson

	for rows.Next() {
		var lesson Lesson

		if err := rows.Scan(&lesson.Id, &lesson.Title, &lesson.Description); err != nil {
			log.Fatal("something went wrong scanning data")
		}

		lessons = append(lessons, lesson)
	}

	for _, lesson := range lessons {
		fmt.Fprintf(w, "%d: %s - %s\n", lesson.Id, lesson.Title, lesson.Description)
	}

	if err := db.Close(); err != nil {
		log.Fatal("something went wrong closing db conn")
	}
}

// getLessonById queries the db for a Lesson matching the id.
func getLessonById(w http.ResponseWriter, r *http.Request) {
	// Declare scany context.
	ctx := context.Background()
	connString := createConnString()
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("the db connection failed")
	}

	var lesson Lesson
	vals := mux.Vars(r)

	query := fmt.Sprintf(`SELECT id, title, description fROM lessons WHERE id = %s`, vals["id"])
	if err := sqlscan.Get(ctx, db, &lesson, query); err != nil {
		log.Fatal("failed to query db")
	}

	if err := db.Close(); err != nil {
		log.Fatal("something went wrong closing the db conn")
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(lesson); err != nil {
		log.Fatal("error serializing lesson data to json")
	}
}

// generateNewLesson deserializes a new Lesson from the request body and returns it to the caller.
func generateNewLesson(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
		return
	}

	var lesson Lesson
	json.NewDecoder(r.Body).Decode(&lesson)

	fmt.Fprintf(w, "Received lesson of Id: %d, Title: %s, Description: %s", lesson.Id, lesson.Title, lesson.Description)
}

func main() {
	router := mux.NewRouter()

	// Register handlers
	router.HandleFunc("/lessons", getLessons).Methods("GET")
	router.HandleFunc("/lessons", generateNewLesson).Methods("POST")
	router.HandleFunc("/lessons/{id}", getLessonById).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", router))
}
