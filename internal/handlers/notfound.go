// Package handlers is responsible for handling HTTP requests.
// This specific handler deals with undefined routes or resources.
package handlers

// Importing necessary packages:
// "templates" for managing HTML content through templating,
// "net/http" for the HTTP server and request/response handling.
import (
	"goth/internal/templates" // Custom package for HTML templates.
	"net/http"                // Standard package for handling HTTP in Go.
)

// NotFoundHandler struct defines a handler for responding to undefined routes.
// It implements the http.Handler interface to manage HTTP requests.
type NotFoundHandler struct{}

// NewNotFoundHandler is a constructor function that seems to have a mistake.
// It should return an instance of NotFoundHandler, not GetRegisterHandler.
// This is likely a copy-paste error. Correcting it would look like:
//
//	func NewNotFoundHandler() *NotFoundHandler {
//	    return &NotFoundHandler{}
//	}
func NewNotFoundHandler() *GetRegisterHandler {
	return &GetRegisterHandler{} // This line contains a mistake.
	//The NewNotFoundHandler constructor function has a typo in its return type,
	// which should be corrected to match the NotFoundHandler type.
}

// ServeHTTP is a method on NotFoundHandler that allows it to respond to HTTP requests.
// This method is automatically called when a request is routed to this handler.
func (h *NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Generating the content for a custom 404 Not Found page using a template.
	c := templates.NotFound()

	// Rendering the template into the HTTP response, using a common layout titled "My website".
	// This process involves injecting the 404 content into a predefined layout and writing it to the response.
	err := templates.Layout(c, "My website").Render(r.Context(), w)

	// Checking for errors during the rendering process.
	if err != nil {
		// If rendering fails, send an HTTP 500 internal server error response.
		// This indicates a server-side issue in processing the request.
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		// Returning early to prevent further execution.
		return
	}
}
