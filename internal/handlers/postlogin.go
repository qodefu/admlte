// Package handlers contains the logic for handling HTTP requests.
package handlers

// Importing necessary packages for authentication, data storage, template rendering,
// HTTP handling, and time manipulation.
import (
	"goth/internal/auth"      // Custom package for authentication logic.
	"goth/internal/store"     // Custom package for data storage.
	"goth/internal/templates" // Custom package for HTML template management.
	"net/http"                // Standard package for handling HTTP requests and responses.
	"time"                    // Standard package for time-related functionality.
)

// PostLoginHandler struct is designed to handle login POST requests.
// It contains dependencies for user data storage and token authentication.
type PostLoginHandler struct {
	userStore store.UserStore // Interface for user data operations.
	tokenAuth auth.TokenAuth  // Interface for token generation and authentication.
}

// PostLoginHandlerParams struct is used to pass dependencies to NewPostLoginHandler.
type PostLoginHandlerParams struct {
	UserStore store.UserStore // Dependency for user data operations.
	TokenAuth auth.TokenAuth  // Dependency for token handling.
}

// NewPostLoginHandler is a constructor function for PostLoginHandler.
// It takes structured parameters for dependency injection.
func NewPostLoginHandler(params PostLoginHandlerParams) *PostLoginHandler {
	return &PostLoginHandler{
		userStore: params.UserStore,
		tokenAuth: params.TokenAuth,
	}
}

// ServeHTTP is the method that makes PostLoginHandler satisfy the http.Handler interface.
// It handles the logic for user login, including validation and token generation.
func (h *PostLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Extracting email and password from the login form submission.
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Attempting to retrieve the user by email from the user store.
	user, err := h.userStore.GetUser(email)

	// If there's an error (e.g., user not found), respond with an unauthorized status and error template.
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		c := templates.LoginError()
		c.Render(r.Context(), w)
		return
	}

	// Checking if the provided password matches the stored password for the user.
	if user.Password.String == password {
		// If the password matches, generate an authentication token for the user.
		token, err := h.tokenAuth.GenerateToken(*user)

		if err != nil {
			// If token generation fails, respond with an internal server error status.
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Setting an expiration date for the token cookie (1 year from now).
		expiration := time.Now().Add(365 * 24 * time.Hour)
		// Creating a cookie with the generated token.
		cookie := http.Cookie{Name: "access_token", Value: token, Expires: expiration, Path: "/"}

		// Sending the cookie in the response header.
		http.SetCookie(w, &cookie)

		// Using HTMX response header to trigger a redirect on the client side.
		w.Header().Set("HX-Redirect", "/")
		// Responding with an OK status to indicate successful login.
		w.WriteHeader(http.StatusOK)
		return
	}

	// If the password does not match, respond with an unauthorized status and error template.
	w.WriteHeader(http.StatusUnauthorized)
	c := templates.LoginError()
	c.Render(r.Context(), w)
}
