package main

import (
	"log"
	"net/http"
)

func main() {
	userService := NewUserService()

	// Start a background routine to periodically clean up inactive users
	go CleanUpInactiveUsers(userService)

	// Initialize routes and start the server
	r := NewRouter(userService)
	port := ":8080"
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
