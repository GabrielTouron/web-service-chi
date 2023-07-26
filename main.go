package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"example/chi/postgresql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		panic(err)
	}

	queries := postgresql.New(db)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello World!"))
	})

	r.Get("/commands", func(w http.ResponseWriter, r *http.Request) {
		commands, err := queries.ListCommands(ctx)
		if err != nil {
			panic(err)
		}

		log.Println(commands)

		jsonData, err := json.Marshal(commands)

		if err != nil {
			fmt.Printf("could not marshal json: %s\n", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})
	http.ListenAndServe(":3000", r)
}
