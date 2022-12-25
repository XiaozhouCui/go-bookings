package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config, can be accessed by any parts of the app
type AppConfig struct {
	UseCache      bool                          // app preference, turn cache on/off
	TemplateCache map[string]*template.Template // cached templates
	InfoLog       *log.Logger                   // pointer to log.Logger
	ErrorLog      *log.Logger                   // pointer to log.Logger
	InProduction  bool                          // identifies env
	Session       *scs.SessionManager           // pointer to scs.SessionManager
}
