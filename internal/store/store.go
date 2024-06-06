// Package store is designed to define data structures and interfaces for storing and retrieving data.
// In this context, it focuses on user-related data operations.
package store

import (
	"goth/internal/store/models"
	"time"
)

// User struct defines the structure for user information within the application.
// It includes basic attributes like email and password.

// UserStore interface declares the operations that can be performed on user data.
// This abstraction allows for different implementations of user data storage,
// such as in-memory stores, database-backed stores, or even external services.
type UserStore interface {
	// CreateUser is a method to add a new user with the given email and password.
	// It should return an error if the user cannot be created, for example, if a user with the same email already exists.
	CreateUser(name, email string, password string) error

	// GetUser is a method to retrieve a user by their email.
	// It returns a pointer to a User struct and an error.
	// The error should be non-nil if the user cannot be found or if there's another issue retrieving the user.
	GetUserByEmail(email string) (models.User, error)
	GetUserById(id int64) (models.User, error)
	DeleteUser(id int64) error
	UpdateUser(name, email, password string, id int64) error
	ListUsers(offset, limit int) []models.User
	GetUserCount() int64
}

type ApptStore interface {
	ListAppts() []models.ListApptRow
	CreateAppt(id int64, apptTime time.Time, status, note string) (models.Appointment, error)
	UpdateAppt(id int64, clientId int64, apptTime time.Time, status, note string) error
	GetApptById(id int64) (models.GetAppointmentRow, error)
}

type ClientStore interface {
	ListClients() []models.Client
}
