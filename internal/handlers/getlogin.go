// Package handlers contains HTTP handlers for the web application.
// It's where the logic for responding to specific routes is defined.
package handlers

// Importing necessary packages:
// "templates" for managing HTML templates,
// "net/http" for handling HTTP requests and responses.
import (
	"goth/internal/templates" // Custom package for template rendering.
	"net/http"                // Standard Go package for HTTP server implementation.
)

// GetLoginHandler is a struct that serves as a handler for the login page.
// It implements the http.Handler interface by having a ServeHTTP method.
type GetLoginHandler struct{}

// NewGetLoginHandler is a constructor function that initializes
// and returns a new instance of GetLoginHandler.
// This pattern is common in Go for creating instances of structs.
func NewGetLoginHandler() *GetLoginHandler {
	return &GetLoginHandler{}
}

// ServeHTTP is a method on GetLoginHandler that allows it to respond to HTTP requests.
// It's called whenever a request is routed to this handler, specifically for rendering the login page.
func (h *GetLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Preparing content for the login page using a function from the templates package.
	// "Login" is passed as a parameter, possibly to set the page title or header.
	c := templates.Login("Login")

	// Rendering the template content into the HTTP response.
	// "My website" might be used as a global title or heading for the layout.
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	// Checking for errors during the rendering process.
	if err != nil {
		// If an error occurs, an HTTP 500 internal server error response is sent back to the client,
		// indicating a server-side problem with rendering the template.
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		// Exiting the function early to prevent further execution after the error.
		return
	}
}
