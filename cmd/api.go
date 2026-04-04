package main

import (
	"log"
	"net/http"
	"time"

	repo "github.com/Ej0416/go-note-app/internal/adapters/postgresql/sqlc"
	"github.com/Ej0416/go-note-app/internal/auth"
	"github.com/Ej0416/go-note-app/internal/env"
	mw "github.com/Ej0416/go-note-app/internal/middleware"
	"github.com/Ej0416/go-note-app/internal/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	config config
	db     *pgxpool.Pool
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

var jwtSecret = []byte(env.GetString("JWT_SECRET", "te-mp-or-ar-ry-ke-y!"))

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // important for rate limiting, analytics and 	tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	authService := auth.NewService(repo.New(app.db), string(jwtSecret))
	authHandler := auth.NewHandler(authService)
	authMiddelware := mw.Auth(jwtSecret)

	// users
	usersService := user.NewService(repo.New(app.db))
	usersHandler := user.NewHandler(usersService)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("server running"))
		})

		// auth
		r.Post("/user/register", authHandler.RegisterUser)
		r.Post("/user/login", authHandler.LoginUser)

		// protected routes
		r.Group(func(r chi.Router) {
			r.Use(authMiddelware)
			r.Get("/user/list",usersHandler.ListUsers)
			r.Get("/user/{id}",usersHandler.GetUserByID)
		})
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at addr %s", app.config.addr)

	return srv.ListenAndServe()
}
