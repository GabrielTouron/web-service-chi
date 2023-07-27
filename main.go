package main

import (
	"example/chi/auth"
	"example/chi/db"
	"example/chi/postgresql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db.Connect()

	queries := postgresql.New(db.Database)

	ap := db.AppResource{Queries: queries}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Mount("/users", usersResource{AppResource: ap}.Routes())
	r.Mount("/", auth.AuthResource{AppResource: ap}.Routes())

	http.ListenAndServe(":3000", r)
}
