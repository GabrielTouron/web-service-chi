package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"example/chi/postgresql"
)

var Database *sql.DB

type AppResource struct {
	Queries *postgresql.Queries
}

func Connect() {
	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	// assign *sql.DB to the global variable
	Database = db
}

// Queries is a helper function to create a new AppResource
func Queries(queries *postgresql.Queries) *AppResource {
	return &AppResource{Queries: queries}
}
