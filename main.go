package main

import (
	"database/sql"
	"fmt"
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
	id          int
	title       string
	description string
}

func createConnString() string {
	user := os.Getenv("DATABASE_USER")
	pw := os.Getenv("DATABASE_PASSWORD")
	return fmt.Sprintf("user=%s port=%s password=%s dbname=%s sslmode=%s", user, port, pw, dbName, sslMode)
}

func GetLessons(w http.ResponseWriter, r *http.Request) {
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

		if err := rows.Scan(&lesson.id, &lesson.title, &lesson.description); err != nil {
			log.Fatal("something went wrong scanning data")
		}

		lessons = append(lessons, lesson)
	}

	for _, lesson := range lessons {
		fmt.Fprintf(w, "%d: %s - %s\n", lesson.id, lesson.title, lesson.description)
	}

	db.Close()
}

func main() {
	router := mux.NewRouter()

	// Register handlers
	router.HandleFunc("/lessons", GetLessons).Methods("GET")

	http.ListenAndServe(":80", router)
}
