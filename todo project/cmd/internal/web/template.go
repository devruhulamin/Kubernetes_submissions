package web

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"
	"time"
)

//go:embed "html" "static"
var Files embed.FS

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
	"now":       time.Now,
}

type TemplateData struct {
	CurrentYear int
	Form        any
	Todos       []string
	Flash       string
	CSRFToken   string
}

func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(Files, "html/*.tmpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		fmt.Print(page)
		name := filepath.Base(page)
		patterns := []string{
			"html/base.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil
}
