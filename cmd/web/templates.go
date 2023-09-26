package main

import (
	"html/template"
	"path/filepath"
	"time"
	"uwDavid/snippetbox/pkg/forms"
	"uwDavid/snippetbox/pkg/models"
)

type templateData struct {
	CurrentYear int
	Flash       string
	Form        *forms.Form
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// initialize template.FuncMap obj => store it in a global var
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// We want to cache templates
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	//use filepath.Glob() to get a slice of all files with *.page.tmpl
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// extract filename
		name := filepath.Base(page)

		// template.New() create new empty template set,
		// then Funcs() to register template.FuncMap
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// add 'layout' to template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		// add 'partial's to template set
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
	// next initialize this cache in main.go
}
