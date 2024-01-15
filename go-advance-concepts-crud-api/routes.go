// routes.go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router with user-related routes
func NewRouter(userService *UserService) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", CreateUserHandler(userService)).Methods("POST")
	r.HandleFunc("/users/{id}", GetUserHandler(userService)).Methods("GET")
	r.HandleFunc("/users", GetUsersHandler(userService)).Methods("GET")
	r.HandleFunc("/users/{id}", DeactivateUserHandler(userService)).Methods("DELETE")

	return r
}

// CreateUserHandler handles user creation requests
func CreateUserHandler(userService *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newUserRequest struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}

		if err := json.NewDecoder(r.Body).Decode(&newUserRequest); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user, err := userService.CreateUser(newUserRequest.Username, newUserRequest.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// GetUserHandler handles requests to retrieve a specific user
func GetUserHandler(userService *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userID := params["id"]

		user, err := userService.GetUser(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// GetUsersHandler handles requests to retrieve all users
func GetUsersHandler(userService *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := userService.GetUsers()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

// DeactivateUserHandler handles requests to deactivate a user
func DeactivateUserHandler(userService *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		userID := params["id"]

		if err := userService.DeactivateUser(userID); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
