// Package middleware is intended to contain middleware functions for HTTP request processing.
package middleware

// Importing necessary packages for context management, cryptographic operations,
// encoding, formatting, and HTTP handling.
import (
	"context"      // For modifying request context.
	"crypto/rand"  // For generating cryptographically secure random strings.
	"encoding/hex" // For encoding binary data into hexadecimal.
	"fmt"          // For formatting strings.
	"net/http"     // For handling HTTP requests and responses.
)

// generateRandomString creates a random string of a specified length.
// It is used for generating nonces for Content-Security-Policy (CSP) headers.
func generateRandomString(length int) string {
	bytes := make([]byte, length) // Allocate a slice of bytes.
	_, err := rand.Read(bytes)    // Fill the slice with random data.
	if err != nil {
		return "" // Return an empty string in case of error.
	}
	return hex.EncodeToString(bytes) // Return the hex-encoded string representation of the random data.
}

// CSPMiddleware returns a new HTTP handler that adds a Content-Security-Policy header.
// This header enhances security by restricting the sources of content that the browser will execute or render.
func CSPMiddleware(next http.Handler) http.Handler {
	// Generate nonces for embedding in CSP header and HTML templates.
	htmxNonce := generateRandomString(16)
	responseTargetsNonse := generateRandomString(16)
	twNonce := generateRandomString(16)

	// Preparing context with nonces to pass to the next handlers.
	// Note: This approach of generating nonces and setting them in context at middleware initialization
	// will result in the same nonce values for all requests, which is not secure.
	// Nonces should be unique per request. This code should be adjusted to generate nonces per request.
	ctx := context.WithValue(context.Background(), "htmxNonce", htmxNonce)
	ctx = context.WithValue(ctx, "twNonce", twNonce)
	ctx = context.WithValue(ctx, "responseTargetsNonse", responseTargetsNonse)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Define the CSP header value, including the generated nonces and a hash for HTMX's injected CSS.
		htmxCSSHash := "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg="
		cspHeader := fmt.Sprintf("default-src 'self'; script-src 'nonce-%s' 'nonce-%s'; style-src 'nonce-%s' '%s';", htmxNonce, responseTargetsNonse, twNonce, htmxCSSHash)

		// Set the CSP header on the response.
		w.Header().Set("Content-Security-Policy", cspHeader)

		// Proceed with the next handler in the chain, passing along the modified request context.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TextHTMLMiddleware returns a new HTTP handler that sets the Content-Type header to text/html.
// This is useful for ensuring that responses are correctly interpreted as HTML by the client.
func TextHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Explicitly set the Content-Type header to indicate the response contains HTML content.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// Proceed with the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
