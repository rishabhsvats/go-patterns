package main

import (
	"log"
	"os"
	"testing"

	"github.com/rishabhsvats/go-patterns/go-breeders/models"
)

var testApp application

func TestMain(m *testing.M) {
	dsn := "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false"
	db, err := initMySQLDB(dsn)
	if err != nil {
		log.Panic(err)
	}
	testApp = application{
		DB:     db,
		Models: *models.New(db),
	}

	os.Exit(m.Run())
}
