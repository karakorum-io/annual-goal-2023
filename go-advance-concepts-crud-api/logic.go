// logic.go
package main

import (
	"errors"
	"sync"
	"time"
)

// UserService handles user-related logic
type UserService struct {
	users     map[string]*User
	usersLock sync.RWMutex
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{
		users: make(map[string]*User),
	}
}

// CreateUser creates a new user and adds it to the database
func (us *UserService) CreateUser(username, email string) (*User, error) {
	user := NewUser(username, email)

	us.usersLock.Lock()
	defer us.usersLock.Unlock()

	if _, exists := us.users[user.ID]; exists {
		return nil, errors.New("user with the same ID already exists")
	}

	us.users[user.ID] = user
	return user, nil
}

// GetUser retrieves a user by ID from the database
func (us *UserService) GetUser(userID string) (*User, error) {
	us.usersLock.RLock()
	defer us.usersLock.RUnlock()

	user, exists := us.users[userID]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// GetUsers returns a list of all users in the database
func (us *UserService) GetUsers() []*User {
	us.usersLock.RLock()
	defer us.usersLock.RUnlock()

	users := make([]*User, 0, len(us.users))
	for _, user := range us.users {
		users = append(users, user)
	}

	return users
}

// DeactivateUser deactivates a user by setting the 'Active' field to false
func (us *UserService) DeactivateUser(userID string) error {
	us.usersLock.Lock()
	defer us.usersLock.Unlock()

	user, exists := us.users[userID]
	if !exists {
		return errors.New("user not found")
	}

	user.Active = false
	return nil
}

// CleanUpInactiveUsers periodically deactivates users who have been inactive for a certain duration
func CleanUpInactiveUsers(us *UserService) {
	for {
		time.Sleep(5 * time.Minute) // Check every 5 minutes (adjust as needed)

		us.usersLock.Lock()
		for _, user := range us.users {
			if !user.Active && time.Since(user.CreatedAt) > 30*time.Minute {
				delete(us.users, user.ID)
			}
		}
		us.usersLock.Unlock()
	}
}
