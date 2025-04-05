// internal/handlers/user_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eamonnk418/goth-stack/internal/service"
	"github.com/eamonnk418/goth-stack/internal/templates/pages"
	"github.com/eamonnk418/goth-stack/internal/utils"
	"github.com/go-chi/chi/v5"
)

// UserHandler handles HTTP requests for user operations.
type UserHandler struct {
	Service service.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{Service: s}
}

// CreateUserAPIHandler handles POST /users to create a new user.
func (h *UserHandler) CreateUserAPIHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.Service.CreateUser(r.Context(), req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUserViewModelHandler handles POST /users to create a new user.
func (h *UserHandler) CreateUserViewModelHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.Service.CreateUser(r.Context(), req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := pages.UserDetails("User", utils.PtrToVal(user))
	component.Render(r.Context(), w)
}

// GetUserAPIHandler handles GET /users/{id} to fetch a user.
func (h *UserHandler) GetUserAPIHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetUserViewModelHandler handles GET /users/{id} to fetch a user.
func (h *UserHandler) GetUserViewModelHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, err := h.Service.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	component := pages.UserDetails("User", utils.PtrToVal(user))
	component.Render(r.Context(), w)
}

// ListUsersAPIHandler handles GET /users to list all users.
func (h *UserHandler) ListUsersAPIHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// ListUsersViewModelHandler handles GET /users to list all users.
func (h *UserHandler) ListUsersViewModelHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := pages.UserList("User", utils.PtrToSliceVal(users))
	component.Render(r.Context(), w)
}
