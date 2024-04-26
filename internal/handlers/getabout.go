// Package handlers is responsible for handling HTTP requests in the application.
package handlers

// Import necessary packages: templates for rendering HTML content,
// and net/http for handling HTTP requests and responses.
import (
	"goth/internal/templates" // Custom package for managing templates.
	"net/http"                // Standard package for HTTP server implementation.
)

// AboutHandLer struct defines a handler for the about page.
// It does not contain any fields but serves as a placeholder for methods.
type AboutHandLer struct{}

// NewAboutHandler is a constructor function that initializes
// and returns a new instance of AboutHandLer.
func NewAboutHandler() *AboutHandLer {
	return &AboutHandLer{}
}

// ServeHTTP is a method on AboutHandLer that satisfies the http.Handler interface,
// making AboutHandLer capable of handling HTTP requests.
// This method is called whenever an HTTP request is routed to this handler.
func (h *AboutHandLer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Generate the content for the about page using the About function from the templates package.
	c := templates.About()

	// Render the layout with the about page content using the Layout function from the templates package.
	// "My website" is passed as the layout title. Render writes the output to the http.ResponseWriter.
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	// Check if there was an error rendering the template.
	if err != nil {
		// If there's an error, send an HTTP 500 internal server error response,
		// indicating a problem with rendering the template.
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
