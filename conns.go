package main

import (
	"fmt"
	"os"
)

const dbName string = "theodinproject_development"
const port string = "5432"
const sslMode string = "disable"

// CreateConnString combines package-level variables and environment vars to build a postgresql conn string.
func CreateConnString() string {
	user := os.Getenv("DATABASE_USER")
	pw := os.Getenv("DATABASE_PASSWORD")
	return fmt.Sprintf("user=%s port=%s password=%s dbname=%s sslmode=%s", user, port, pw, dbName, sslMode)
}
