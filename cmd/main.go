package main

// Import necessary packages for routing, context management, error handling, logging, HTTP server functionality, etc.
import (
	"context"
	"errors"
	"fmt"
	"goth/internal/auth/tokenauth" // Package for JWT token authentication logic.
	"goth/internal/handlers"       // Handlers for HTTP routes.
	"goth/internal/store/dbstore"  // Data store for user information.
	"goth/internal/templates"
	"goth/internal/validator"
	v "goth/internal/validator"
	"io"
	"log/slog" // Structured logging.
	"net/http"
	"os"        // Access to operating system functionality like signals and stdout.
	"os/signal" // Signal handling for graceful shutdown.
	"syscall"
	"time"

	m "goth/internal/middleware" // Custom middleware package.

	"github.com/go-chi/chi/v5"            // Chi router for handling HTTP requests.
	"github.com/go-chi/chi/v5/middleware" // Built-in middleware from Chi.
	"github.com/go-chi/jwtauth/v5"        // JWT authentication middleware.
)

// TokenFromCookie extracts the JWT from the access_token cookie.
func TokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func main() {
	// Initialize structured JSON logging.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Initialize the Chi router.
	r := chi.NewRouter()

	// Setup the user store and token authentication with a secret key.
	userStore := dbstore.NewUserStore()
	tokenAuth := tokenauth.NewTokenAuth(tokenauth.NewTokenAuthParams{
		SecretKey: []byte("secret"),
	})

	// Setup static file serving for the /static route.
	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Define routes and middleware.
	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,    // Log every request.
			m.TextHTMLMiddleware, // Set Content-Type headers to text/html.
			m.CSPMiddleware,      // Apply Content Security Policy headers.
			m.UtilMiddleware,
			jwtauth.Verify(tokenAuth.JWTAuth, TokenFromCookie), // Verify JWT tokens from cookies.
		)

		// Define route handlers for not found, home, about, registration, and login.
		r.NotFound(handlers.NewNotFoundHandler().ServeHTTP)
		// Other routes omitted for brevity.
		r.Get("/", handlers.NewHomeHandler().ServeHTTP)

		r.Route("/admin", func(r chi.Router) {
			r.Get("/dashboard", func(w http.ResponseWriter, r *http.Request) {
				// id := chi.URLParam(r, "id")
				templates.Layout(templates.DashContent(), "Smart 1").Render(r.Context(), w)
			})
			r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
				// id := chi.URLParam(r, "id")
				users := userStore.ListUsers()
				templates.Layout(templates.UserContent(users), "Smart 1").Render(r.Context(), w)
			})
			r.Get("/users/hx/launchModal", func(w http.ResponseWriter, r *http.Request) {
				// id := chi.URLParam(r, "id")
				w.Header().Set("HX-Trigger", "show-global-modal-form")
				io.WriteString(w, "Click!!!")
			})
			r.Post("/users/hx/addNew", func(w http.ResponseWriter, r *http.Request) {
				// first validate
				// validate fail, return form with error
				nameVal := r.FormValue("name")
				emailVal := r.FormValue("email")
				pwdVal := r.FormValue("password")
				pwdConfirm := r.FormValue("passwordConfirmation")

				validations := templates.UserValidations{
					Name:                 v.New("name", nameVal, v.NotEmpty("Name")),
					Email:                v.New("email", emailVal, v.NotEmpty("Email"), v.EmailFmt),
					Password:             v.New("password", pwdVal, v.NotEmpty("Password")),
					PasswordConfirmation: v.New("passwordConfirmation", pwdConfirm, v.NotEmpty("PasswordConfirmation"), v.PasswordMatch(pwdVal)),
				}
				validator.ValidateFields(&validations)

				// validation pass, create user, return empty form
				if validator.ValidationOk(&validations) {
					userStore.CreateUser(nameVal, emailVal, pwdVal)
					w.Header().Set("HX-Trigger", `{"close-global-modal-form": [{"foo": 1, "message": "User Added", "tags": "Success!"}]}`)
				}
				templates.UserForm(validations).Render(r.Context(), w)
			})
			r.Get("/users/hx/list", func(w http.ResponseWriter, r *http.Request) {
				users := userStore.ListUsers()

				templates.UserTable(users).Render(r.Context(), w)

			})
		})
		r.Get("/about", handlers.NewAboutHandler().ServeHTTP)

		r.Get("/register", handlers.NewGetRegisterHandler().ServeHTTP)

		r.Post("/register", handlers.NewPostRegisterHandler(handlers.PostRegisterHandlerParams{
			UserStore: userStore,
		}).ServeHTTP)

		r.Get("/login", handlers.NewGetLoginHandler().ServeHTTP)

		r.Post("/login", handlers.NewPostLoginHandler(handlers.PostLoginHandlerParams{
			UserStore: userStore,
			TokenAuth: tokenAuth,
		}).ServeHTTP)
	})

	// Setup channel and signal notification for graceful shutdown.
	killSig := make(chan os.Signal, 1)
	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	// Define the server address and create the HTTP server.
	port := ":8080"
	srv := &http.Server{
		Addr:    port,
		Handler: r,
	}

	// Start the server in a goroutine to allow for graceful shutdown handling.
	go func() {
		// ListenAndServe blocks until the server is stopped.
		err := srv.ListenAndServe()
		// Handle server closure and errors.
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Log server start and wait for a shutdown signal.
	logger.Info("Server started", slog.String("port", port))
	<-killSig // Block until a shutdown signal is received.

	// Begin shutdown process.
	logger.Info("Shutting down server")

	// Create a context with a timeout for server shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server.
	if err := srv.Shutdown(ctx); err != nil {
		// Log any errors encountered during shutdown.
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("Server shutdown complete")
}

/*
Key Points:

    Structured Logging: Uses slog for JSON formatted logs, enhancing observability and
	debugging capabilities.

    Chi Router: Utilizes the Chi router for defining HTTP routes and middleware, offering a
	lightweight and idiomatic way to build web applications in Go.

    JWT Authentication: Integrates JWT-based authentication for secure access to certain routes.

    Graceful Shutdown: Implements signal handling for graceful server shutdown, ensuring

	connections are properly closed without abruptly terminating ongoing requests.

    Middleware and Handlers: Demonstrates how to organize and use middleware and route handlers,
	including custom logic for token extraction and response content types.

    Static File Serving: Serves static files from a designated directory, making it easy to serve
	assets like images, JavaScript, and CSS files.

This main file orchestrates the setup and lifecycle of the web application, emphasizing security, maintainability, and user experience.
*/
