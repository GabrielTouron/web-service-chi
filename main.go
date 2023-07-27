package main

import (
	"context"
	"encoding/json"
	"example/chi/auth"
	"example/chi/db"
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

	db.Connect()

	queries := postgresql.New(db.Database)

	ap := db.AppResource{Queries: queries}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/commands", func(w http.ResponseWriter, r *http.Request) {

		commands, err := ap.Queries.ListCommands(ctx)
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

	r.Mount("/users", usersResource{ap}.Routes())
	r.Mount("/", auth.AuthResource{ap}.Routes())

	http.ListenAndServe(":3000", r)
}
