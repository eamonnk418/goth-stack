// cmd/main.go
package main

import (
	"log"
	"net/http"

	"github.com/eamonnk418/goth-stack/internal/auth"
	"github.com/eamonnk418/goth-stack/internal/config"
	"github.com/eamonnk418/goth-stack/internal/handlers"
	"github.com/eamonnk418/goth-stack/internal/service"
	"github.com/eamonnk418/goth-stack/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load the config for our app
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// we have the config but we need to share its state across the entire app

	// Initialize the store (data access layer)
	userStore := store.NewInMemoryUserStore()

	// Initialize the service (business logic layer) with the store dependency injected
	userService := service.NewUserService(userStore)

	// Initialize the handlers (HTTP layer) with the service injected
	userHandler := handlers.NewUserHandler(userService)

	// Setup Auth handler
	authInstance := auth.NewAuth(cfg)
	authHandler := &handlers.AuthHandler{Auth: authInstance}

	// Set up Chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Auth routes group
	r.Route("/auth/github", func(r chi.Router) {
		r.Get("/", authHandler.Login)
		r.Get("/callback", authHandler.Callback)
	})

	// API routes group
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))
		r.Get("/users/{id}", userHandler.GetUserAPIHandler)
		r.Get("/users", userHandler.ListUsersAPIHandler)
	})

	// Web routes group
	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.HomeHandler{}.ServeHTTP)
		// Define routes
		r.Post("/users", userHandler.CreateUserViewModelHandler)
		r.Get("/users/{id}", userHandler.GetUserViewModelHandler)
		r.Get("/users", userHandler.ListUsersViewModelHandler)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)
}
