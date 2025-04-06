// internal/handlers/auth_handler.go
package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/eamonnk418/goth-stack/internal/auth"
)

// AuthHandler handles GitHub OAuth login.
type AuthHandler struct {
	Auth *auth.Auth
}

// Login redirects the user to GitHub's OAuth2 consent page.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// In production, generate and store a unique state per session.
	state := "randomStateString"
	url := h.Auth.AuthCodeURL(state)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Callback handles GitHub's OAuth2 callback.
func (h *AuthHandler) Callback(w http.ResponseWriter, r *http.Request) {
	// Validate the state parameter (omitted for brevity).
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	// Exchange the code for an access token.
	token, err := h.Auth.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Code exchange failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Log the token for demonstration. In practice, use this token to create a client and fetch user data.
	log.Printf("Access Token: %s", token.AccessToken)

	// Redirect the user back to the application (home page).
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
