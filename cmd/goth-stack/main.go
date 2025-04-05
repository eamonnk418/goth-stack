// cmd/main.go
package main

import (
	"log"
	"net/http"

	"github.com/eamonnk418/goth-stack/internal/handlers"
	"github.com/eamonnk418/goth-stack/internal/service"
	"github.com/eamonnk418/goth-stack/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Initialize the store (data access layer)
	userStore := store.NewInMemoryUserStore()

	// Initialize the service (business logic layer) with the store dependency injected
	userService := service.NewUserService(userStore)

	// Initialize the handlers (HTTP layer) with the service injected
	userHandler := handlers.NewUserHandler(userService)

	// Set up Chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// API routes group
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Get("/users/{id}", userHandler.GetUserAPIHandler)
		r.Get("/users", userHandler.ListUsersAPIHandler)
	})

	// Web routes group
	r.Route("/", func(r chi.Router) {
		// Define routes
		r.Post("/users", userHandler.CreateUserViewModelHandler)
		r.Get("/users/{id}", userHandler.GetUserViewModelHandler)
		r.Get("/users", userHandler.ListUsersViewModelHandler)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
