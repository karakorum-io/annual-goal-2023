// model.go
package main

import "time"

// User represents a user entity
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Active    bool      `json:"active"`
}

// NewUser creates a new user with the given username and email
func NewUser(username, email string) *User {
	return &User{
		ID:        GenerateID(),
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
		Active:    true,
	}
}

// GenerateID generates a unique ID for the user
func GenerateID() string {
	// Implementation to generate a unique ID (e.g., UUID)
	// For simplicity, we'll use a placeholder
	return "user-id-placeholder"
}
