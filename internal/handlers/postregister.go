// Package handlers includes the logic for processing HTTP requests.
package handlers

// Importing necessary packages:
// "store" for data persistence operations,
// "templates" for rendering HTML templates,
// "net/http" for handling HTTP requests and responses.
import (
	"goth/internal/store"     // Custom package for data storage operations.
	"goth/internal/templates" // Custom package for managing HTML templates.
	"net/http"                // Standard package for HTTP server functionalities.
)

// PostRegisterHandler struct is responsible for handling registration POST requests.
// It contains a userStore which is an interface for user data operations.
type PostRegisterHandler struct {
	userStore store.UserStore // Interface to interact with the user data storage.
}

// PostRegisterHandlerParams struct for passing dependencies to NewPostRegisterHandler.
type PostRegisterHandlerParams struct {
	UserStore store.UserStore // Dependency for user data operations.
}

// NewPostRegisterHandler is a constructor function that initializes a PostRegisterHandler.
// It takes a struct of parameters for dependency injection, promoting decoupling and testability.
func NewPostRegisterHandler(params PostRegisterHandlerParams) *PostRegisterHandler {
	return &PostRegisterHandler{
		userStore: params.UserStore,
	}
}

// ServeHTTP makes PostRegisterHandler satisfy the http.Handler interface,
// enabling it to respond to registration HTTP requests.
func (h *PostRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Extracting email and password from the registration form.
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Attempting to create a new user with the provided email and password.
	err := h.userStore.CreateUser(name, email, password)

	// If there is an error during user creation (e.g., user already exists), respond with an error.
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// If user creation is successful, render the registration success template.
	c := templates.RegisterSucces() // Note: There's a typo here; it should likely be RegisterSuccess().
	err = c.Render(r.Context(), w)

	// Check for errors during the rendering of the success template.
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}

	// After successful registration and template rendering, execution ends.
	// The user is presumably informed of their successful registration via the rendered template.
}
