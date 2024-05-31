package tokenauth

// Import necessary packages.
import (
	"errors" // For creating error messages.
	// Presumably, your internal package for user management.
	"goth/internal/store/models"
	"time" // For setting token expiration times.

	"github.com/go-chi/jwtauth/v5" // JWT middleware for the Chi router.
	"github.com/golang-jwt/jwt"    // The JWT library for Go.
)

// TokenAuth struct holds the configuration for JWT authentication.
type TokenAuth struct {
	JWTAuth   *jwtauth.JWTAuth // Pointer to a JWTAuth instance for creating and validating tokens.
	secretKey []byte           // The secret key used for signing tokens.
}

// NewTokenAuthParams struct to pass parameters when creating a new TokenAuth instance.
type NewTokenAuthParams struct {
	SecretKey []byte // Secret key for JWT signing.
}

// Constructor function for TokenAuth. It initializes a new JWTAuth instance with HS256 signing method.
func NewTokenAuth(params NewTokenAuthParams) *TokenAuth {
	jwtAuth := jwtauth.New("HS256", []byte(params.SecretKey), nil) // HS256 is a commonly used signing method.
	return &TokenAuth{
		JWTAuth:   jwtAuth,
		secretKey: params.SecretKey,
	}
}

// GenerateToken creates a new JWT for a user. It includes the user's email and an expiration time.
func (a *TokenAuth) GenerateToken(user models.User) (string, error) {
	// Define the payload for the token. 'exp' sets the expiration time.
	payload := map[string]interface{}{
		"email": user.Email.String,                     // Include the user's email in the token payload.
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours.
	}

	// Encode the payload into a token string.
	_, tokenString, err := a.JWTAuth.Encode(payload)
	if err != nil {
		return "", err // Return an error if token encoding fails.
	}

	return tokenString, nil // Return the token string.
}

// ValidateToken checks the validity of a token string and returns the claims if valid.
func (a *TokenAuth) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Initialize an empty MapClaims object.
	claims := jwt.MapClaims{}

	// Parse the token string into a token object, validating it using the secret key.
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return a.secretKey, nil // Use the stored secret key for validation.
	})

	if err != nil {
		return nil, err // Return an error if token parsing fails.
	}

	if !token.Valid {
		return nil, errors.New("invalid token") // Return an error if the token is not valid.
	}

	return claims, nil // Return the claims from the token if it is valid.
}

/*
Key Points:

    Interface Definition: The TokenAuth interface defines a contract for JWT authentication mechanisms,
	ensuring consistency and interchangeability of authentication implementations in your application.

    Dependency on store.User: The GenerateToken method's dependency on the store.User type implies
	that the JWT generation process likely includes user-specific information, such as user ID or email,
	within the token's claims.

    Use of jwt.MapClaims: The ValidateToken method's return type, jwt.MapClaims, allows for a
	flexible representation of the JWT's payload, accommodating various claim types and structures
	without requiring a predefined struct.

    Separation of Concerns: By defining this interface in the auth package, you separate the
	concerns of token management from other parts of your application, adhering to good software
	design principles. Implementations of this interface can then be used throughout your
	application to handle authentication-related tasks, promoting code reusability and
	maintainability.

This interface acts as a foundation for implementing JWT authentication in your Go applications,
allowing for different concrete implementations that fulfill the interface's contract. Whether
using an in-memory store, a database, or an external authentication service, any implementation
of TokenAuth will provide a consistent way to generate and validate JWTs.
*/
