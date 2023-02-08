/*
	 _____ _  _ _   _ _  _ ___  ___ ___  ___  ___  ___ _____
	|_   _| || | | | | \| |   \| __| _ \/ __|/ _ \| __|_   _|
	  | | | __ | |_| | .` | |) | _||   /\__ \ (_) | _|  | |
	  |_| |_||_|\___/|_|\_|___/|___|_|_\|___/\___/|_|   |_|

1.	Package name: 							render
2.	Date of creation:						2023-01-20
3.	Name of creator of module:				Benoit Frontzak
4.	History of modification:				N/A (v1)
5.	Summary of what the module does:		Render template with customized functions & Create template cache
6.	Functions in that module:				RenderTemplate; CreateTemplateCache
7.	Variables accessed by the module:		w http.ResponseWriter, tmpl string (wanted template), app *config.AppConfig

Important: All the templates must follow the naming convention as follow:
Template:	name.page.gohtml	Where name is the template name that you can choose, while '.page.gohtml' is required !
Layout:		name.layout.gohtml	Where name is the layout name that you can choose, while '.layout.gohtml' is required !
*/
package render

import (
	"bytes"
	"fmt"
	"frontend/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// myTemplates holds the path to find our templates
var myTemplates = "./templates/"

// tmplFunctions holds a map (name, func) for customized
// functions that we use with gohtml templates
var tmplFunctions = template.FuncMap{
	"inc": func(x int) int {
		return x + 1
	},
	"dec": func(x int) int {
		return x - 1
	},
	// "rowID": func(x primitive.ObjectID) string {
	// 	return string(x.Hex())
	// },
	"mkSlice": func(args ...string) []string {
		return args
	},
	"isInclude": func(s string, a []string) bool {
		inc := false
		for _, e := range a {
			if e == s {
				inc = true
			}
		}
		return inc
	},
}

// app holds the configuration of the app
var app *config.AppConfig

// NewTemplate set the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

// RenderTemplate pull out a specific template out of templates cache and execute it
func RenderTemplate(w http.ResponseWriter, tmpl string, data any) {
	var tc map[string]*template.Template
	var err error
	if app.UseCache {
		// Get the template cache from the app config
		tc = app.TemplateCache
	} else {
		// Build the template cache
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("could not create the template cache")
		}
	}
	// Pull the wanted template (tmpl) out of the map
	// exit program is can't find the template requested
	t, ok := tc[tmpl]
	if !ok {
		log.Fatalf("could not get %s from template cache", tmpl)
	}

	// Create a bytes buffer to store the template in memory
	buf := new(bytes.Buffer)

	// Execute the template and store the value into the buffer
	// Exit program is can't excute template
	err = t.Execute(buf, data)
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer to ResponseWriter
	// ignore the number of bytes written
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	// My Cache holds all the name of our templates with his associated (ready to use) template
	mc := make(map[string]*template.Template)

	// Fetch all available template pages
	pages, err := filepath.Glob(myTemplates + "*.page.gohtml")
	if err != nil {
		return mc, err
	}

	// Loop over all pages to populate mc (my cache)
	for _, page := range pages {
		// Get only the filename (not the fullpath)
		name := filepath.Base(page)

		// Template Set (with customized functions)
		ts, err := template.New(name).Funcs(tmplFunctions).ParseFiles(page)
		if err != nil {
			return mc, err
		}
		// Does this template matches a layout
		matches, err := filepath.Glob(myTemplates + "*.layout.gohtml")
		if err != nil {
			return mc, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(myTemplates + "*.layout.gohtml")
			if err != nil {
				return mc, err
			}
		}
		// Populate mc
		mc[name] = ts
	}

	return mc, nil
}
