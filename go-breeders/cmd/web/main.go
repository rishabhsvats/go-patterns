package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/rishabhsvats/go-patterns/go-breeders/adapters"
	"github.com/rishabhsvats/go-patterns/go-breeders/configuration"
)

const PORT = ":3000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {

	app := application{
		templateMap: make(map[string]*template.Template),
	}
	flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false", "DSN")
	flag.Parse()

	//get database

	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}
	//jsonBackend := &adapters.JSONBackend{}
	//jsonAdapter := &adapters.RemoteService{Remote: jsonBackend}

	xmlBackend := &adapters.XMLBackend{}
	xmlAdapter := &adapters.RemoteService{Remote: xmlBackend}

	app.App = configuration.New(db, xmlAdapter)

	srv := &http.Server{
		Addr:              PORT,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web application on port: ", PORT)
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
