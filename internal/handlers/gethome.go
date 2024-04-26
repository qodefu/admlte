// Package handlers is designated for handling HTTP requests within the application.
package handlers

// Importing necessary packages.
// "templates" for handling HTML templates,
// "net/http" for web server functionalities,
// "github.com/go-chi/jwtauth/v5" for JWT authentication.
import (
	"goth/internal/templates" // Custom package for managing HTML templates.
	"net/http"                // Standard library for HTTP server and client.

	"github.com/go-chi/jwtauth/v5" // JWT middleware for Go, used for authentication.
)

// HomeHandler struct defines a handler for the homepage.
// It does not contain fields but will implement http.Handler by having ServeHTTP method.
type HomeHandler struct{}

// NewHomeHandler is a constructor function that creates
// and returns a new instance of HomeHandler.
func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

// ServeHTTP is implemented by HomeHandler to satisfy the http.Handler interface,
// allowing it to respond to HTTP requests.
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve the JWT token and claims from the request context.
	// Ignoring the token itself and error for simplicity, focusing on claims.
	_, claims, _ := jwtauth.FromContext(r.Context())

	// Attempt to extract the "email" claim as a string from the JWT claims.
	email, ok := claims["email"].(string)

	// Check if the "email" claim was successfully retrieved.
	if !ok {
		// If not ok (meaning the user is not authenticated or the claim is missing),
		// serve the guest version of the homepage.
		c := templates.GuestIndex()
		err := templates.Layout(c, "My website").Render(r.Context(), w)

		// Check for errors in rendering the template.
		if err != nil {
			// If an error occurs, send an HTTP 500 response.
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}

		// Return early for unauthenticated users.
		return
	}

	// If the email claim is present, render the authenticated version of the homepage.
	c := templates.Index(email)
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	// Again, check for errors in rendering the template.
	if err != nil {
		// Send an HTTP 500 response if an error occurs during rendering.
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
