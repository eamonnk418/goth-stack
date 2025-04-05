package handlers

import (
	"net/http"

	"github.com/eamonnk418/goth-stack/internal/templates/pages"
)

// HomeHandler serves the index page.
type HomeHandler struct{}

// ServeHTTP implements the http.Handler interface for HomeHandler.
func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	homePage := pages.HomePage(pages.Props{
		Title: "Homepage",
	})
	homePage.Render(r.Context(), w)
}
