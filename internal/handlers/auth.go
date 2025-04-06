package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

// AuthHandler manages authentication routes.
type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Callback handles the OAuth callback from the provider.
func (h *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	// Store the provider in the context for gothic
	ctx := context.WithValue(r.Context(), "provider", provider)
	r = r.WithContext(ctx)

	_, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("Authentication failed: %v", err), http.StatusInternalServerError)
		return
	}

	log.Printf("Cookies on callback: %+v\n", r.Cookies())

	// Here, you can create a session for the user
	// For example, store user information in a session or database
	// fmt.Fprintln(w, user)

	// Redirect to the home page or dashboard after successful login
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// Login initiates the authentication process.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	// Store the provider in the context for gothic
	ctx := context.WithValue(r.Context(), "provider", provider)
	r = r.WithContext(ctx)

	gothic.BeginAuthHandler(w, r)
}

// Logout logs the user out and redirects to the home page.
func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	// Store the provider in the context for gothic
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))

	if err := gothic.Logout(w, r); err != nil {
		http.Error(w, fmt.Sprintf("Logout failed: %v", err), http.StatusInternalServerError)
		return
	}

	// Redirect to the home page or login page after logout
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
