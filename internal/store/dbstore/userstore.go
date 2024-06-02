package dbstore

import (
	// Standard package for creating error objects.

	"context"
	"goth/internal/store" // Internal package where the User struct is defined.
	"goth/internal/store/models"

	"github.com/jackc/pgx/v5/pgtype"
)

// UserStore struct holds a slice of users, acting as an in-memory storage mechanism.
// This is a simple approach and not suitable for production environments due to its non-persistent nature.
type UserStore struct {
	db *models.Queries
}

// NewUserStore initializes and returns a new instance of UserStore.
// It pre-populates the store with a default user for demonstration or testing purposes.
func NewUserStore(q *models.Queries) *UserStore {
	return &UserStore{
		db: q,
	}
}

func (s *UserStore) UpdateUser(name, email, password string, id int64) error {
	return s.db.UpdateUser(context.Background(), models.UpdateUserParams{
		ID:       id,
		Name:     name,
		Email:    pgtype.Text{String: email, Valid: true},
		Password: pgtype.Text{String: password, Valid: true},
	})
}

func (s *UserStore) DeleteUser(id int64) error {
	return s.db.DeleteUser(context.Background(), id)
}

// CreateUser attempts to add a new user to the store.
// It checks for existing users with the same email to avoid duplicates and returns an error if found.
func (s *UserStore) CreateUser(name, email, password string) error {
	_, err := s.db.CreateUser(context.Background(), models.CreateUserParams{
		Name:     name,
		Email:    pgtype.Text{String: email, Valid: true},
		Password: pgtype.Text{String: password, Valid: true},
	})
	return err
}

// GetUser searches for a user by email and returns the user object if found.
// If no user is found with the provided email, it returns an error.
func (s *UserStore) GetUser(email string) (models.User, error) {
	return s.db.GetUserByEmail(context.Background(), pgtype.Text{String: email, Valid: true})
}

// GetUser is a method to retrieve a user by their email.
// It returns a pointer to a User struct and an error.
// The error should be non-nil if the user cannot be found or if there's another issue retrieving the user.
func (s UserStore) GetUserByEmail(email string) (models.User, error) {
	return s.db.GetUserByEmail(context.Background(), pgtype.Text{String: email, Valid: true})
}

func (s UserStore) GetUserById(id int64) (models.User, error) {
	return s.db.GetUser(context.Background(), id)
}

func (s *UserStore) ListUsers(offset, limit int) []models.User {
	ret, _ := s.db.ListUsers(context.Background(), models.ListUsersParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	return ret
}

func (s UserStore) GetUserCount() int64 {
	ret, _ := s.db.GetUserCount(context.Background())
	return ret
}

type UserPagination struct {
	store.AbstractPagination[models.User]
	store store.UserStore
}

func NewUserPagination(url string, repo store.UserStore, pg int) UserPagination {
	super := store.MkAbsPgtor[models.User](5, pg, url)
	ret := UserPagination{
		super,
		repo,
	}
	ret.Child = ret
	return ret
}

func (thing UserPagination) Items() []models.User {
	return thing.store.ListUsers((thing.CurPage-1)*thing.ItemsPerPage, thing.ItemsPerPage)
}

func (thing UserPagination) Total() int {
	return int(thing.store.GetUserCount())
}
