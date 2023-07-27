package auth

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"example/chi/db"
	"net/http"
)

type AuthResource struct {
	AppResource db.AppResource
}

type loginPayload struct {
	Email    string 
	Password string 
}

func (ar AuthResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/login", ar.Login)

	return r
}

func (ar AuthResource) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queries := ar.AppResource.Queries

	var login loginPayload

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(login.Password)

	user, err := queries.GetUserByEmail(context.Background(), login.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	match, err := ComparePasswordAndHash(login.Password, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !match {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}
