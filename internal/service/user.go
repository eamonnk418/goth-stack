// internal/service/user_service.go
package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/eamonnk418/goth-stack/internal/store"
)

// UserService defines business logic for managing users.
type UserService interface {
	CreateUser(ctx context.Context, name, email string) (*store.User, error)
	GetUser(ctx context.Context, id int) (*store.User, error)
	ListUsers(ctx context.Context) ([]*store.User, error)
}

type userServiceImpl struct {
	store store.UserStore
}

// NewUserService injects a UserStore into the service.
func NewUserService(s store.UserStore) UserService {
	return &userServiceImpl{
		store: s,
	}
}

func (svc *userServiceImpl) CreateUser(ctx context.Context, name, email string) (*store.User, error) {
	// Simple business validation
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}
	if !containsAt(email) {
		return nil, errors.New("invalid email address")
	}

	user := &store.User{
		Name:  name,
		Email: email,
	}
	if err := svc.store.Save(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}
	return user, nil
}

func (svc *userServiceImpl) GetUser(ctx context.Context, id int) (*store.User, error) {
	return svc.store.GetByID(ctx, id)
}

func (svc userServiceImpl) ListUsers(ctx context.Context) ([]*store.User, error) {
	return svc.store.GetAll(ctx)
}

// containsAt is a simple helper to check if an email contains '@'.
func containsAt(s string) bool {
	for _, ch := range s {
		if ch == '@' {
			return true
		}
	}
	return false
}
