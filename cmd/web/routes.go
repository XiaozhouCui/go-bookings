package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// create a multiplexer (mux)
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // use Recoverer middleware
	mux.Use(NoSurf)               // add CSRF protection
	mux.Use(SessionLoad)          // load and save session on every request

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
