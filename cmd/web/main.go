package main

import (
	"bookings/internal/config"
	"bookings/internal/driver"
	"bookings/internal/handlers"
	"bookings/internal/helpers"
	"bookings/internal/models"
	"bookings/internal/render"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

// package level variable
const portNumber = ":8080"

// global app config variable
var app config.AppConfig

// pointer to scs.SessionManager
var session *scs.SessionManager

// pointer to log.Logger
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		// stop the app if there is an error
		log.Fatal(err)
	}
	defer db.SQL.Close()

	// print in terminal
	fmt.Printf("Starting application on port %s\n", portNumber)

	// create a server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	app.InProduction = false

	// add loggers to the app
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // give information about the error
	app.ErrorLog = errorLog

	// initialise session
	session = scs.New()
	session.Lifetime = 24 * time.Hour              // keep session for 24 hours
	session.Cookie.Persist = true                  // keep session after browser close
	session.Cookie.SameSite = http.SameSiteLaxMode // default SameSite attribute
	session.Cookie.Secure = app.InProduction       // encrypted (https)

	// save session into global app config
	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=54321 dbname=bookings user=postgres password=postgres")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")

	// create a map (tc) to cache all templates
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return nil, err
	}

	// store the template cache to the app config
	app.TemplateCache = tc
	app.UseCache = false

	// declare repository variable
	repo := handlers.NewRepo(&app, db)
	// pass the repo back to the handlers
	handlers.NewHandlers(repo)
	// give render package access to the app config
	render.NewTemplates(&app) // reference to app using pointer
	// give helper function access to the app config
	helpers.NewHelpers(&app)

	return db, nil
}
