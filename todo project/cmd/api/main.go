package main

import (
	"fmt"
	"log"
	"net/http"
	"ruhultodo/cmd/internal/env"
	"time"
)

func main() {
	app := application{
		cfg: config{
			port: env.GetString("PORT", "8080"),
		},
	}

	mux := http.NewServeMux()

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
