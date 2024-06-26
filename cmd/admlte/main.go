package main

// Import necessary packages for routing, context management, error handling, logging, HTTP server functionality, etc.
import (
	"context"
	"errors"
	"fmt"
	"goth/internal/auth/tokenauth" // Package for JWT token authentication logic.
	"goth/internal/components"
	"goth/internal/utils"
	"log"

	// Package for JWT token authentication logic.
	"goth/internal/config"
	"goth/internal/handlers" // Handlers for HTTP routes.
	appts "goth/internal/handlers/admin/appointments"
	"goth/internal/handlers/admin/users"
	"goth/internal/store/dbstore"

	"goth/internal/store/models"
	"goth/internal/templates"
	"goth/internal/templates/admin"
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
	"github.com/jackc/pgx/v5/pgxpool"
)

// TokenFromCookie extracts the JWT from the access_token cookie.
func TokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return ""
	}
	return cookie.Value
}

type hfunc func(req m.RequestScope) error

func wrapH(handler hfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := handler(m.ReqScope(r.Context()))
		if err != nil {
			slog.Error("handler error", err)
			w.WriteHeader(500)
			w.Write([]byte("internal error"))
		}

	}
}

// type CreateComponent func(req m.RequestScope) components.RComp

var compRegistry = make(map[string]components.CreateComp)

var idGen = utils.NewIdGen("ComponentsRegistry")

func registerComponent(regComp components.RegComp) {
	// var compId = idGen.Id(compName)
	// var factory = regFunc(compId)
	compRegistry[regComp.Id] = regComp.Factory
}

func handleComponent(req m.RequestScope) error {
	//get componentId
	if req.IsComponentReq() {
		return compRegistry[req.ComponentId()](req).Render()
	}

	return errors.New("not component request")
}

func main() {
	// Initialize structured JSON logging.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := chi.NewRouter()

	// Setup the user store and token authentication with a secret key.
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", "app1", "app1", "lovelace", 9000, "app1")

	conn, err := pgxpool.New(context.Background(), connectStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	queries := models.New(conn)
	userStore := dbstore.NewUserStore(queries)
	apptStore := dbstore.NewApptStore(queries)
	cliStore := dbstore.NewCliStore(queries)
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

		cfgRoutes := config.Routes()

		r.Get("/rpc-dispatch", wrapH(func(req m.RequestScope) error {
			return handleComponent(req)
		}))
		r.Delete("/rpc-dispatch", wrapH(func(req m.RequestScope) error {
			return handleComponent(req)
		}))
		r.Post("/rpc-dispatch", wrapH(func(req m.RequestScope) error {
			return handleComponent(req)
		}))
		// admin users
		r.Group(func(r chi.Router) {
			listUsersHandler := users.NewListUsersHandler(userStore)
			r.Get(cfgRoutes.Admin.Users.Base, func(w http.ResponseWriter, r *http.Request) {
				paginator := dbstore.NewUserPagination("", "", userStore, 1)
				templates.Layout(users.UserContent(paginator), "Smart 1").Render(r.Context(), w)
			})
			r.Get(cfgRoutes.Admin.Users.HX.AddUserModal, listUsersHandler.HxAddUserModal)
			r.Get(cfgRoutes.Admin.Users.HX.EditUserModal, listUsersHandler.HxEditUserModal)
			r.Delete(cfgRoutes.Admin.Users.HX.DeleteUserModal, listUsersHandler.HxDeleteUserModal)
			r.Post(cfgRoutes.Admin.Users.HX.Create, listUsersHandler.HxCreateUser)
			r.Post(cfgRoutes.Admin.Users.HX.Update, listUsersHandler.HxUpdateUser)
			r.Delete(cfgRoutes.Admin.Users.HX.Delete, listUsersHandler.HxDeleteUser)
			r.Get(cfgRoutes.Admin.Users.HX.List, wrapH(listUsersHandler.HxListUsers))
		})

		// Appointments
		r.Group(func(r chi.Router) {
			apptsHandler := appts.NewApptsHandler(apptStore, cliStore)
			registerComponent(apptsHandler.ApptListComp())

			r.Get(cfgRoutes.Admin.Appt.Base, func(w http.ResponseWriter, r *http.Request) {
				rcomp := apptsHandler.Component(appts.COMP_LIST_APPT)(m.ReqScope(r.Context()))
				templates.Layout(components.WrapRPC(rcomp), "Appointment").Render(r.Context(), w)
			})

			r.Get(cfgRoutes.Admin.Appt.Base2, func(w http.ResponseWriter, r *http.Request) {

				pgtor := dbstore.NewApptPagination(apptStore, 1)
				templates.Layout(appts.ApptContent(pgtor), "Appointment").Render(r.Context(), w)
			})
			r.Get(cfgRoutes.Admin.Appt.Create, handlers.Func(apptsHandler.CreateForm))
			r.Post(cfgRoutes.Admin.Appt.SaveNew, wrapH(apptsHandler.SaveNew))
			r.Get(cfgRoutes.Admin.Appt.UpdateAppt, handlers.Func(apptsHandler.UpdateAppt))
			r.Get(cfgRoutes.Admin.Appt.HX.List, wrapH(apptsHandler.ListAppt))
			r.Get(cfgRoutes.Admin.Appt.HX.DeleteApptConfirm, wrapH(apptsHandler.DeleteApptConfirm))
			r.Delete(cfgRoutes.Admin.Appt.HX.DeleteAppt, wrapH(apptsHandler.DeleteAppt))
		})

		// Dashboard
		r.Get(cfgRoutes.Admin.Dashboard.Base, func(w http.ResponseWriter, r *http.Request) {
			// id := chi.URLParam(r, "id")
			templates.Layout(admin.DashContent(), "Smart 1").Render(r.Context(), w)
		})

		// r.Get("/about", handlers.NewAboutHandler().ServeHTTP)

		// r.Get("/register", handlers.NewGetRegisterHandler().ServeHTTP)

		// r.Post("/register", handlers.NewPostRegisterHandler(handlers.PostRegisterHandlerParams{
		// 	UserStore: userStore,
		// }).ServeHTTP)

		// r.Get("/login", handlers.NewGetLoginHandler().ServeHTTP)

		// r.Post("/login", handlers.NewPostLoginHandler(handlers.PostLoginHandlerParams{
		// 	UserStore: userStore,
		// 	TokenAuth: tokenAuth,
		// }).ServeHTTP)
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
