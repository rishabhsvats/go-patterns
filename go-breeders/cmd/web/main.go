package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

const PORT = ":3000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
}

type appConfig struct {
	useCache bool
}

func main() {

	app := application{
		templateMap: make(map[string]*template.Template),
	}
	flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
	flag.Parse()

	srv := &http.Server{
		Addr:              PORT,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web application on port: ", PORT)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
