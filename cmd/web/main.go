package main

import (
	"bookings/internal/config"
	"bookings/internal/handlers"
	"bookings/internal/models"
	"bookings/internal/render"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

// package level variable
const portNumber = ":8080"

// global app config variable
var app config.AppConfig

// pointer to scs.SessionManager
var session *scs.SessionManager

func main() {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	// initialise session
	session = scs.New()
	session.Lifetime = 24 * time.Hour              // keep session for 24 hours
	session.Cookie.Persist = true                  // keep session after browser close
	session.Cookie.SameSite = http.SameSiteLaxMode // default SameSite attribute
	session.Cookie.Secure = app.InProduction       // encrypted (https)

	// save session into global app config
	app.Session = session

	// create a map (tc) to cache all templates
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// store the template cache to the app config
	app.TemplateCache = tc
	app.UseCache = false

	// declare repository variable
	repo := handlers.NewRepo(&app)
	// pass the repo back to the handlers
	handlers.NewHandlers(repo)
	// give render package access to the app config
	render.NewTemplates(&app) // reference to app using pointer

	// print in terminal
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// create a server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
