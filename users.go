package main

import (
	"context"
	"encoding/json"
	"example/chi/auth"
	"example/chi/db"
	"example/chi/postgresql"
	"log"

	"github.com/go-chi/chi/v5"
	"net/http"
)

type usersResource struct {
	AppResource db.AppResource
}

type auth2params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	keyLength   uint32
	saltLength  int
}

type userResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (rs usersResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", rs.List)
	r.Post("/", rs.Post)

	return r
}

func (rs usersResource) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("users list of stuff.."))
}

func (rs usersResource) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user postgresql.User
	queries := rs.AppResource.Queries

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	password := user.Password

	hash, err := auth.CreateHash(password, auth.DefaultParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// create userParams
	userParams := postgresql.CreateUserParams{
		Email:    user.Email,
		Name:     user.Name,
		Password: hash,
	}

	// create user
	user, err = queries.CreateUser(context.Background(), userParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get user by email named userFound
	userFound, err := queries.GetUserByEmail(context.Background(), user.Email)

	match, err := auth.ComparePasswordAndHash(userFound.Password, hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userResponse := userResponse{
		Name:  user.Name,
		Email: user.Email,
	}

	log.Printf("Password match: %v\n", match)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
}
