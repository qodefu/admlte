package dbstore

import (
	// Standard package for creating error objects.

	"goth/internal/store" // Internal package where the User struct is defined.
	"goth/internal/store/models"
)

// UserStore struct holds a slice of users, acting as an in-memory storage mechanism.
// This is a simple approach and not suitable for production environments due to its non-persistent nature.
type UserStore struct {
}

// NewUserStore initializes and returns a new instance of UserStore.
// It pre-populates the store with a default user for demonstration or testing purposes.
func NewUserStore() *UserStore {
	panic("bad")
}

func (s *UserStore) UpdateUser(name, email string, password string) error {
	panic("bad")
}

func (s *UserStore) DeleteUser(email string) error {
	panic("bad")
	// Appending the new user to the users slice if no duplicate is found.
}

// CreateUser attempts to add a new user to the store.
// It checks for existing users with the same email to avoid duplicates and returns an error if found.
func (s *UserStore) CreateUser(name, email string, password string) error {
	panic("bad")
}

// GetUser searches for a user by email and returns the user object if found.
// If no user is found with the provided email, it returns an error.
func (s *UserStore) GetUser(email string) (*models.User, error) {
	panic("bad")
}

func (s *UserStore) ListUsers() []models.User {
	panic("bad")
}

type UserPagination struct {
	baseUrl      string
	curPage      int
	itemsPerPage int
	store        store.UserStore
}

func NewUserPagination(url string, store store.UserStore, pg int) UserPagination {
	return UserPagination{baseUrl: url,
		curPage:      pg,
		itemsPerPage: 5,
		store:        store,
	}
}

func (thing UserPagination) PageUrl(page int) string {
	panic("bad")
}

func (thing UserPagination) Pages() []int {
	panic("bad")
}

func (thing UserPagination) Items() []models.User {
	panic("bad")
}

func (thing UserPagination) PageCount() int {
	panic("bad")
}

func (thing UserPagination) CurrentPage() int {
	panic("bad")
}

func (thing UserPagination) PerPage() int {
	panic("bad")
}

func (thing UserPagination) PreviousPageUrl() string {
	panic("bad")
}

func (thing UserPagination) Total() int {
	panic("bad")
}

func (thing UserPagination) NextPageUrl() string {
	panic("bad")
}
