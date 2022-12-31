package render

import (
	"bookings/internal/config"
	"bookings/internal/models"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig
var pathToTemplates = "./templates"

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// Template renders a template
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	// declare template cache (tc) as a map
	var tc map[string]*template.Template

	// get preference "UseCache" from app config
	if app.UseCache {
		// get template cache from app config
		tc = app.TemplateCache
	} else {
		// this is just used for testing, so taht we rebuild the cache on every request
		tc, _ = CreateTemplateCache()
	}

	// get requested template (t) from cache
	t, ok := tc[tmpl] // tmpl is the file name (e.g. "home.page.tmpl")
	if !ok {
		// if template not found, return an error
		return errors.New("could not get template from cache")
	}

	buf := new(bytes.Buffer) // create buffer to hold bytes

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td) // write template to buffer, with template data

	// render the template
	_, err := buf.WriteTo(w) // write from buffer to http response
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return err
	}

	return nil
}

// Cache templates to avoid loading template files from disk on every request
// it returns a map (myCache) and an error
func CreateTemplateCache() (map[string]*template.Template, error) {
	// create an empty map to cache templates
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates)) // return a slice of strings (full path)
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page) // return the filename without path
		// parse the file and save it into a template set (ts)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// find all layout-templates
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		// if there are layout-templates, associate them with page-templates (ts)
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates)) // use "=" instead of ":="
			if err != nil {
				return myCache, err
			}
		}

		// at the end of loop, save the template into cache (e.g. myCache['home.page.tmpl'])
		myCache[name] = ts
	}

	return myCache, nil
}
