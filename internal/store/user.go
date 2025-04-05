// internal/store/user_store.go
package store

import (
	"context"
	"errors"
	"sync"
)

// User represents a simple user model.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserStore defines methods for persisting and retrieving users.
type UserStore interface {
	GetByID(ctx context.Context, id int) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Save(ctx context.Context, user *User) error
}

// InMemoryUserStore is a simple in-memory implementation of UserStore.
type InMemoryUserStore struct {
	mu     sync.RWMutex
	users  map[int]*User
	nextID int
}

// NewInMemoryUserStore creates a new InMemoryUserStore.
func NewInMemoryUserStore() UserStore {
	// return &InMemoryUserStore{
	// 	users:  make(map[int]*User),
	// 	nextID: 1,
	// }
	return &InMemoryUserStore{
		users: map[int]*User{
			1: {ID: 1, Name: "Alice", Email: "alice@example.com"},
			2: {ID: 2, Name: "Bob", Email: "bob@example.com"},
		},
	}
}

func (s *InMemoryUserStore) GetByID(ctx context.Context, id int) (*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, ok := s.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *InMemoryUserStore) GetAll(ctx context.Context) ([]*User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		list = append(list, user)
	}
	return list, nil
}

func (s *InMemoryUserStore) Save(ctx context.Context, user *User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// If the user has no ID, assign a new one.
	if user.ID == 0 {
		user.ID = s.nextID
		s.nextID++
	}
	s.users[user.ID] = user
	return nil
}
