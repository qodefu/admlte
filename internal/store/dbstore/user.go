// Package dbstore is designated for data storage operations, specifically for user data in this context.
package dbstore

// Importing necessary packages:
// "errors" for handling error conditions,
// "goth/internal/store" presumably contains shared types or interfaces, including the User struct definition.
import (
	"errors"              // Standard package for creating error objects.
	"goth/internal/store" // Internal package where the User struct is defined.
)

// UserStore struct holds a slice of users, acting as an in-memory storage mechanism.
// This is a simple approach and not suitable for production environments due to its non-persistent nature.
type UserStore struct {
	users []store.User // Slice of User structs to store user data.
}

// NewUserStore initializes and returns a new instance of UserStore.
// It pre-populates the store with a default user for demonstration or testing purposes.
func NewUserStore() *UserStore {
	return &UserStore{
		users: []store.User{
			// Initializing with a default user.
			{
				Name:     "one",
				Email:    "1@example.com",
				Password: "password", // Note: Storing passwords in plain text is insecure and not recommended.
			},
		},
	}
}

// CreateUser attempts to add a new user to the store.
// It checks for existing users with the same email to avoid duplicates and returns an error if found.
func (s *UserStore) CreateUser(name, email string, password string) error {
	for _, user := range s.users {
		if user.Email == email {
			// Preventing duplicate user registration.
			return errors.New("user already exists")
		}
	}

	// Appending the new user to the users slice if no duplicate is found.
	s.users = append(s.users, store.User{Name: name, Email: email, Password: password})
	return nil
}

// GetUser searches for a user by email and returns the user object if found.
// If no user is found with the provided email, it returns an error.
func (s *UserStore) GetUser(email string) (*store.User, error) {
	for _, user := range s.users {
		if user.Email == email {
			// Found the user, return a pointer to the User struct.
			return &user, nil
		}
	}

	// If the loop completes without finding a user, return an error indicating this.
	return nil, errors.New("user not found")
}

func (s *UserStore) ListUsers() []store.User {
	return s.users[:]
}
