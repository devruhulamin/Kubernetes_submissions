package main

import (
	"fmt"
	"log"
	"net/http"
	"ruhultodo/cmd/internal/env"
	"ruhultodo/cmd/internal/web"

	"time"
)

func main() {
	tmplCache, err := web.NewTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app := application{
		cfg: config{
			port: env.GetString("PORT", "8080"),
		},
		templateCache: tmplCache,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", app.home)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", app.cfg.port),
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server started in port %s", app.cfg.port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := &web.TemplateData{
		Todos: []string{},
	}
	app.render(w, http.StatusOK, "home.tmpl", data)
}
