// Declare the package name. This code is part of the 'auth' package,
// which likely contains functionality related to authentication.
package auth

// Import necessary packages.
import (
	"goth/internal/store" // Import your internal store package, presumably for user management.

	"github.com/golang-jwt/jwt" // Import the JWT library for handling JSON Web Tokens.
)

// TokenAuth is an interface that outlines methods for JWT authentication.
// It sets the contract for any struct that implements it, ensuring
// they provide these methods for generating and validating JWT tokens.
type TokenAuth interface {
	// GenerateToken should take a user object (from your internal store package)
	// and return a JWT string and any error that occurred during the token generation.
	// This method is intended to create a new JWT based on user details, typically after
	// a successful login attempt.
	GenerateToken(user store.User) (string, error)

	// ValidateToken should take a JWT string and return the claims contained within
	// the token as a MapClaims object, along with any error that occurred during validation.
	// This method is used to verify the authenticity and validity of a provided JWT,
	// extracting claims like user identity or token expiry as needed.
	ValidateToken(tokenString string) (jwt.MapClaims, error)
}

/*
Key Points:
    JWT Authentication: This module provides a simple interface to generate JWTs for authenticated
	users and to validate JWT tokens in subsequent requests.

    Expiration: Generated tokens include an expiration claim (exp), set to 24 hours from
	generation. This helps mitigate security risks associated with stolen tokens.

    Security: It uses the HS256 signing method, which requires a secret key for token generation
	and validation. The secret key should be kept secure and not exposed in your codebase or
	version control.

    Flexibility: By including user email in the token payload, the application can easily
	identify the user associated with a token. Additional claims can be added as needed for
	your application's requirements.
*/
