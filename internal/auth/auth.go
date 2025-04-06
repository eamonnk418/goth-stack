package auth

import (
	"net/http"

	"github.com/eamonnk418/goth-stack/internal/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

const (
	sessionSecret = "your-session-secret" // Replace with a secure secret in production
	maxAge        = 86400 * 30            // 30 days
	isProd        = false                 // Set to true in production
)

// NewAuth initializes the authentication providers and session store.
func NewAuth(cfg *config.Config) {
	store := sessions.NewCookieStore([]byte(sessionSecret))
	store.MaxAge(maxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd
	store.Options.SameSite = http.SameSiteLaxMode

	gothic.Store = store

	goth.UseProviders(
		github.New(cfg.ClientID, cfg.ClientSecret, cfg.RedirectURL),
		// Add other providers here
	)
}
