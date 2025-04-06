// internal/auth/auth.go
package auth

import (
	"github.com/eamonnk418/goth-stack/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type Auth struct {
	*oauth2.Config
}

// NewAuth creates a new Auth instance configured for GitHub.
func NewAuth(cfg *config.Config) *Auth {
	return &Auth{
		Config: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectURL,
			Scopes:       []string{"user:email"},
			Endpoint:     github.Endpoint,
		},
	}
}
