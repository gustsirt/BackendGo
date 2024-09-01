package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gustsirt/BackendGo/models"
	"github.com/gustsirt/BackendGo/repository"
	"github.com/gustsirt/BackendGo/server"
	"github.com/segmentio/ksuid"
)

const (
	HASH_COST = 8
)

type SignUpResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SignUpLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpLoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), HASH_COST)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// var user = models.User{
		// 	Email:    request.Email,
		// 	Password: string(hashedPassword),
		// 	Id:       id.String(),
		// }
		var user = models.User{
			Email:    request.Email,
			Password: request.Password,
			Id:       id.String(),
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			ID:    user.Id,
			Email: user.Email,
		})

	}
}
