// Package handlers is designated for handling HTTP requests within the application.
// It defines the logic for various endpoints, including the registration page in this instance.
package handlers

// Importing necessary packages:
// "templates" for rendering HTML templates tailored to different parts of the application,
// "net/http" for handling HTTP requests and responses.
import (
	"goth/internal/templates" // Custom package for managing HTML templates.
	"net/http"                // Standard package for HTTP server functionalities.
)

// GetRegisterHandler struct serves as a handler for the registration page.
// It implements the http.Handler interface, enabling it to respond to HTTP requests.
type GetRegisterHandler struct{}

// NewGetRegisterHandler is a constructor function that initializes
// and returns a new instance of GetRegisterHandler.
// This is a common pattern in Go to encapsulate the creation of new instances.
func NewGetRegisterHandler() *GetRegisterHandler {
	return &GetRegisterHandler{}
}

// ServeHTTP is a method on GetRegisterHandler that allows it to handle HTTP requests.
// It's automatically called when a request to the associated route is received.
func (h *GetRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Generating the content for the registration page using the RegisterPage function
	// from the templates package. This likely prepares HTML content specific to registration.
	c := templates.RegisterPage()

	// Attempting to render the prepared content into the HTTP response writer,
	// using a layout template titled "My website". This could imply a common layout
	// or theme is being applied across the site.
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	// Checking for any errors that occurred during the rendering process.
	if err != nil {
		// If an error is encountered, send an HTTP 500 internal server error response,
		// indicating a server-side issue with rendering the template.
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		// Early return to halt further execution in case of error.
		return
	}

}
