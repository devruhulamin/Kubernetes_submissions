package main

import "html/template"

type application struct {
	cfg           config
	templateCache map[string]*template.Template
}

type config struct {
	port string
}
