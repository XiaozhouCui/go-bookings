package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// create a map to cache templates, as loading template files from disk on every request is expensive
var tc = make(map[string]*template.Template)

// RenderTemplateSimple renders templates using html/templates
func RenderTemplateSimple(w http.ResponseWriter, t string) {
	// t is the template file name, e.g. home.page.tmpl
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCacheSimple(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the tempalte in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	// execute the cached template
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCacheSimple(t string) error {
	// t is the template file name, e.g. home.page.tmpl
	// create a slice of string
	templates := []string{
		fmt.Sprintf("./templates/%s", t), // Sprintf returns the string
		"./templates/base.layout.tmpl",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = tmpl

	return nil
}
